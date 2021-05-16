const express = require('express')
const http = require('http');
const cors = require('cors')
const app = express()
const { Etcd3, EtcdLockFailedError } = require('etcd3');
import { startServer, deleteServer, getInstanceList } from "./gcp"

console.log("process.env.ETCD", process.env.ETCD)

const MAX_PEER_PER_SERVER = 100
const MAX_LOAD = 70 //start new instance after 70% load
const MAX_PEER_HOST_MAPPING = {
    "5.9.18.28": 10
}

const client = new Etcd3({
    "hosts": process.env.ETCD || "localhost:2379"
});
const { Server } = require("socket.io");

let avaiable_hosts = {}
let sessions = {}

const getHostData = () => {
    return avaiable_hosts
}

app.use(cors())

app.use(express.json());
app.use(express.urlencoded({
    extended: true
}));

const server = http.createServer(app);
const io = new Server(server);




io.on('connection', (socket) => {
    console.log('a user connected');
    socket.broadcast.emit('data', getHostData());
    socket.on("get", () => {
        socket.broadcast.emit('data', getHostData());
    })
});



app.get('/', async (req, res) => {
    res.send('ping server')
})

let session_host_tree = {}

let timeoutId = false

const debounce_calc_session_stats = async () => {
    if (timeoutId) {
        clearTimeout(timeoutId)
    }
    timeoutId = setTimeout(async () => {
        await calc_session_stats()
    }, 10)
}

const calc_session_stats = async () => {
    const sessions = await client.getAll().prefix('/session/').keys();
    // console.log("sessions", sessions)

    if (sessions.length === 0)
        session_host_tree = {}

    session_host_tree = {}
    sessions.forEach(session => {
        if (session.indexOf("node") > 0 && session.indexOf("peer") > 0 && session.indexOf("track") > 0) {
            const session_split = session.split("/")
            const name = session_split[2]
            const node = session_split[4]
            const peer = session_split[6]
            const track = session_split[8]
            const trackType = session_split[10]
            if (!session_host_tree[node])
                session_host_tree[node] = {}

            if (!session_host_tree[node][name]) {
                session_host_tree[node][name] = {}
            }
            if (!session_host_tree[node][name][peer]) {
                session_host_tree[node][name][peer] = [{
                    "track": track,
                    "trackType": trackType
                }]
            } else {
                session_host_tree[node][name][peer].push({
                    "track": track,
                    "trackType": trackType
                })
            }
        }
        if (session.indexOf("node") > 0 && session.indexOf("peer") > 0) {
            const session_split = session.split("/")
            const name = session_split[2]
            const node = session_split[4]
            const peer = session_split[6]
            if (!session_host_tree[node])
                session_host_tree[node] = {}

            if (!session_host_tree[node][name]) {
                session_host_tree[node][name] = {}
            }
            if (!session_host_tree[node][name][peer]) {
                session_host_tree[node][name][peer] = []
            }
        }
        if (session.indexOf("node") > 0) {
            const session_split = session.split("/")
            const name = session_split[2]
            const node = session_split[4]
            if (!session_host_tree[node])
                session_host_tree[node] = {}

            if (!session_host_tree[node][name])
                session_host_tree[node][name] = {}
        }
    })
    Object.keys(session_host_tree).forEach(host => {
        const rooms = Object.keys(session_host_tree[host]).length
        let peers = 0
        Object.keys(session_host_tree[host]).forEach(room => {
            peers = peers + Object.keys(session_host_tree[host][room]).length
        })
        session_host_tree[host]["rooms"] = rooms
        session_host_tree[host]["peers"] = peers
    })
}

app.get("/stats", async (req, res) => {
    const sessions = await client.getAll().prefix('/session/').keys();
    res.json({
        "tree": session_host_tree,
        "sessions": sessions,
        "hosts": await client.getAll().prefix('available-hosts/').keys()
    })
})


const findAndGetHost = async (id, req, res) => {
    const sessionNode = await client.get("/session/" + id).string()

    // const currentHosts = await client.getAll().prefix('available-hosts/').keys();

    if (Object.keys(avaiable_hosts).length === 0) {
        res.send("NO_HOSTS_RETRY")
        return
    }
    let hostExists = false
    if (sessionNode) {
        //TODO: if server went down, then avaiable hosts will be empty but session will be assigned to a host
        //but what is that host hasn't sent a ping yet so its not there in available hosts
        // also if a sfu crashes, then its keys won't be deleted
        hostExists = Object.keys(avaiable_hosts).find(key => {
            return key.replace("::", ":").replace("available-hosts/", "") === sessionNode.replace("::", ":")
        })
    }
    if (hostExists) {
        console.log("session node", sessionNode)
        res.send(sessionNode.replace("::", ":"))
    } else {
        // this is a problem here because its possible that we are getting very high requests
        // and even before sfu starts we can give another host to the same session
        // need to solve using etcd again

        // get lock for session

        // TODO: it is also possible that hosts die's in between the api response as well
        // need to see how to handle that maybe on client side by quering nodes again?
        try {
            const existingHost = await client.get("/temp" + id)
            if (existingHost) {
                console.log("host just assigned", existingHost)
                return res.send(existingHost.replace("::", ":"))
            }
            await client.lock("/query/" + id).do(() => {
                return new Promise(async (resolve) => {

                    const currentHosts = {
                        ...avaiable_hosts
                    }

                    console.log("current hosts", Object.keys(currentHosts))
                    const sortedKeys = Object.keys(currentHosts).sort((key1, key2) => {
                        console.log(currentHosts[key1])
                        console.log(currentHosts[key2])
                        const cpu1 = parseFloat(currentHosts[key1][0].split("-")[1])
                        const cpu2 = parseFloat(currentHosts[key2][0].split("-")[1])
                        return cpu1 > cpu2
                    })
                    const sortedHostKey = sortedKeys[0]
                    const filterhosts = Object.keys(currentHosts).filter(key => {
                        const cpu1 = parseFloat(currentHosts[key][0].split("-")[1])
                        console.log("cpu", cpu1)
                        return cpu1 < 10
                    })
                    let hostKey = ""
                    if (filterhosts.length > 0) {
                        console.log("idle hosts", filterhosts)
                        hostKey = filterhosts[Math.floor(Math.random() * filterhosts.length)];
                    } else {
                        hostKey = sortedHostKey
                    }
                    const load = currentHosts[hostKey][0]
                    console.log("hostKey", hostKey, load)
                    hostKey = hostKey.replace("available-hosts/", "").replace("::", ":")


                    // there can be a gap in between when host is assigned and sfu starts a connection
                    // in that gap we will assign the same host
                    const lease = client.lease(process.env.LEASE_TEMP_TIMEOUT || 5, {
                        autoKeepAlive: false
                    });
                    lease.on('lost', err => {
                        console.log('We lost our lease as a result of this error:', err);
                    })
                    await lease.put("/temp" + id).value(hostKey).exec();
                    res.send(hostKey)
                    resolve()
                })
            })
        } catch (e) {
            if (e instanceof EtcdLockFailedError) {
                console.log("lock is already aquired by another process")
                setTimeout(() => {
                    findAndGetHost(id, req, res)
                }, 100)
            } else {
                throw e
            }

        }
    }

}

app.get("/session/:id", async (req, res) => {
    const id = req.params.id
    if (id.length === 0) {
        return res.send("Invalid ID")
    }
    // this will provide the host which can serve new client
    // if session already exists in a sfu it will return that
    // if session doesn't exists it will simply provide an sfu in which session exists already

    // first lets check if session already exists


    await findAndGetHost(id, req, res)

})


/** this will check loads acorss all server and add/delete instances as needed */

let waitForHost = false

let gcp_hosts = []
let gcp_hosts_deadmap = {}
let gcp_inactive_map = {}

const autoScaleServerLoads = async () => {
    const currentHosts = {}
    Object.keys(avaiable_hosts).forEach(key => {
        let ip = key.replace("available-hosts/", "")
        if (ip.indexOf(":") !== -1) {
            ip = ip.substr(0, ip.indexOf(":"))
        }
        const max = 5
        let count = 0
        let sum = 0
        avaiable_hosts[key].forEach(val => {
            if (count < max)
                sum = sum + parseFloat(val[0].split("-")[1])
            count = count + 1
        })
        currentHosts[ip] = sum / count
    })
    console.log("autoScaleServerLoads current hosts", currentHosts)


    let gcp_instance_list = await getInstanceList()
    gcp_instance_list = JSON.parse(gcp_instance_list)

    let gcp_ip_name_map = {}
    if (Object.keys(currentHosts).length > 0) {
        //first check if there are any dead gcp server
        console.log("check for any dead instances....")
        gcp_instance_list.forEach(async current_instance => {
            if (current_instance["status"] !== "RUNNING") {
                return
            }
            const host_ip = current_instance["networkInterfaces"][0]["accessConfigs"].find(cfg => cfg.name === "external-nat")["natIP"]
            gcp_ip_name_map[host_ip] = {
                "name": current_instance["name"],
                "zone": current_instance['zone'].split("/").slice(-1).pop()
            }
            if (Object.keys(currentHosts).find(host => host === host_ip)) {
                console.log("host ip found all good")
                if (gcp_hosts_deadmap[host_ip]) {
                    delete gcp_hosts_deadmap[host_ip]
                }
            } else {
                console.log("this looks like a dead gcp instance wait to delete it", host_ip)
                if (!gcp_hosts_deadmap[host_ip])
                    gcp_hosts_deadmap[host_ip] = new Date().getTime()

                const timeDiff = (new Date().getTime() - gcp_hosts_deadmap[host_ip]) / 1000
                console.log("time diff on dead instance", timeDiff)
                if (timeDiff > (process.env.GCP_DEAD_HOST_DELETE_WAIT || 5 * 60)) {
                    console.log("instance is dead since", timeDiff, "so deleating it!")
                    await deleteServer(gcp_ip_name_map[host_ip]["name"], gcp_ip_name_map[host_ip]["zone"])
                }
            }
        })
    }

    let skipProcess = false
    let filterhosts = Object.keys(currentHosts).filter(host => {
        // if (process.env.MY_IP && host === process.env.MY_IP) {
        //     console.log("my ip", process.env.MY_IP, "host", host)
        //     //TEMP code skipping current server for load
        //     return false
        // }
        const cpu1 = parseFloat(currentHosts[host])
        console.log("cpu", cpu1, "host", host)
        return cpu1 < MAX_LOAD
    })

    if (Object.keys(currentHosts).length < (process.env.MINIMUM_HOSTS || 1)) {
        filterhosts = []
    }
    if (filterhosts.length > 0) {
        console.log("all good server loads under 70%", filterhosts)
    } else {
        if (Object.keys(currentHosts).length < (process.env.MINIMUM_HOSTS || 1)) {
            console.log("need minimum no of hosts so starting")
        } else {
            console.log("all server load over 70% need to start a new server")
        }

        if (waitForHost) {
            console.log("waiting for server to start")

            gcp_hosts.map(host => {

                if (!host["ready"]) {
                    const found = Object.keys(currentHosts).find(host => {
                        console.log("checking ", host, " with", host["host_ip"])
                        return host.indexOf(host["host_ip"]) !== -1
                    })
                    host["ready"] = found
                    if (found)
                        waitForHost = false
                }
                return host
            })

        } else {
            const current_instance = await startServer()
            if (current_instance) {
                waitForHost = true
                // console.log("current instance", current_instance)
                const host_ip = current_instance["networkInterfaces"][0]["accessConfigs"].find(cfg => cfg.name === "external-nat")["natIP"]
                console.log("host ip", host_ip)
                gcp_hosts.push({
                    "host_ip": host_ip,
                    "name": current_instance["name"],
                    "zone": current_instance["zone"].split("/").slice(-1).pop(),
                    "ready": false
                })
            } else {
                console.log("instance creating failed!")
            }
        }
        skipProcess = true
        //not going further
    }
    if (!skipProcess) {
        await calc_session_stats()

        Object.keys(currentHosts).find(async host => {
            if (process.env.MY_IP && host === process.env.MY_IP) {
                console.log("skipping current host", host)
            } else {
                if (!gcp_ip_name_map[host]) {
                    console.log("this is not a gcp host so not looking at deleting...", host)
                }
                if (Object.keys(session_host_tree).find(key => key.indexOf(host) !== -1)) {
                    console.log("session active on ", host)
                } else {
                    console.log("no sessions active on host ", host, " can be deleted!", gcp_ip_name_map[host])
                    if (!gcp_inactive_map[host]) {
                        gcp_inactive_map[host] = new Date().getTime()
                    }
                    const timeDiff = (new Date().getTime() - gcp_inactive_map[host]) / 1000
                    if (timeDiff > (process.env.GCP_EMPTY_SESSION_DELETE_WAIT || 1 * 60)) {
                        console.log("host has had no session for 1 min so deleting it now!")
                        if (Object.keys(currentHosts).length > (process.env.MINIMUM_HOSTS || 1)) {
                            await deleteServer(gcp_ip_name_map[host]["name"], gcp_ip_name_map[host]["zone"])
                            await new Promise(r => setTimeout(() => { r() }, 10000))
                            return true //so that it doesn't delete more
                        } else {
                            console.log("cannot delete need atleast one host")
                        }

                    } else {
                        console.log("not deleting host as timeDiff less than 1 min", timeDiff)
                    }
                }
            }

            return false
        })
    }
    setTimeout(async () => {
        await autoScaleServerLoads()
    }, process.env.AUTO_SCALE_TIMEOUT || 5000)
}

const PORT = process.env.PORT || 4000
server.listen(PORT, async function () {
    // server ready to accept connections here
    console.log("server has started", PORT)

    setTimeout(async () => {
        await autoScaleServerLoads()
        //waiting for hosts to ping for atleast 5sec
    }, 5000)


    const hosts = await client.getAll().prefix('available-hosts/').keys();
    console.log('available hosts:', hosts);

    const sessions = await client.getAll().prefix('/session/').keys();
    console.log('available sessions:', sessions);
    // sessions.forEach(async session => await client.delete().key(session)) //temp

    client.watch().prefix("/session/").create().then(watcher => {
        watcher
            .on('delete', async (res) => {
                const session = res.key.toString()
                if (session.indexOf("node") === -1) {
                    console.log("session closed", session)
                    const sessions = await client.getAll().prefix('/session/' + session).keys();
                    console.log('deleting all session keys', sessions);
                    await Promise.all(sessions.map(session => {
                        return new Promise(async (resolve) => {
                            await client.delete().key(session)
                            resolve()
                        })
                    }))
                }
                await debounce_calc_session_stats()
            })
            .on('put', async (res) => {
                // const session = res.key.toString()
                // const data = res.value.toString()
                // console.log("session created", session, data)
                await debounce_calc_session_stats()
            });
    });

    client.watch().prefix("available-hosts/").create().then(watcher => {
        watcher
            .on('delete', async (res) => {
                const host = res.key.toString()
                console.log("delete", host)
                if (avaiable_hosts[host]) {
                    delete avaiable_hosts[host]
                }

                //if host gets dropped remove all sessions from the host
                const sessions = await client.getAll().prefix('/session/').keys();
                console.log(host, 'deleting available sessions:', sessions);

                const actualSession = {}
                await Promise.all(sessions.map(session => {
                    return new Promise(async (resolve) => {
                        // console.log(session.replace("::", ":"), host.replace("::", ":").replace("available-hosts/", ""), "index of", session.replace("::", ":").indexOf(host.replace("::", ":").replace("available-hosts/", "")))
                        if (session.replace("::", ":").indexOf(host.replace("::", ":").replace("available-hosts/", "")) !== -1) {
                            await client.delete().key(session)
                            actualSession[session.split("/")[2]] = true
                        }
                        resolve()
                    })

                }))

                console.log("also delete", Object.keys(actualSession))
                Object.keys(actualSession).forEach(async session => {
                    await client.delete().key("/session/" + session)
                })


            })
            .on('put', async (res) => {
                const host = res.key.toString()
                const data = res.value.toString()
                console.log(host, ' ping ', data) //, data

                io.emit("pingdata", {
                    "host": host,
                    "data": data
                })
                if (avaiable_hosts[host]) {

                    //keep maximum 30 records per host

                    if (avaiable_hosts[host].length >= 30) {
                        avaiable_hosts[host].pop()
                    }
                    avaiable_hosts[host].unshift(data)
                } else {
                    avaiable_hosts[host] = [data]
                }
            });
    });

});

