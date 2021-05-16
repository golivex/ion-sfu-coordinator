const express = require('express')
const http = require('http');
const cors = require('cors')
const app = express()
const { Etcd3 } = require('etcd3');

console.log("process.env.ETCD", process.env.ETCD)

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



app.get('/', (req, res) => {
    res.send('ping server')
})


app.get("/stats", async (req, res) => {
    const sessions = await client.getAll().prefix('/session/').keys();
    res.json(sessions)
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

                    console.log("current hosts", currentHosts)
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
                    console.log("hostKey", hostKey, currentHosts[hostKey][0])
                    hostkey = hostKey.replace("available-hosts/", "").replace("::", ":")

                    // there can be a gap in between when host is assigned and sfu starts a connection
                    // in that gap we will assign the same host
                    const lease = client.lease(10, {
                        autoKeepAlive: false
                    });
                    lease.on('lost', err => {
                        console.log('We lost our lease as a result of this error:', err);
                    })
                    await lease.put("/temp" + id).value(hostkey).exec();
                    res.send(hostkey)
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



const PORT = process.env.PORT || 4000
server.listen(PORT, async function () {
    // server ready to accept connections here
    console.log("server has started", PORT)

    const hosts = await client.getAll().prefix('available-hosts/').keys();
    console.log('available hosts:', hosts);

    const sessions = await client.getAll().prefix('/session/').keys();
    console.log('available sessions:', sessions);
    // sessions.forEach(async session => await client.delete().key(session))

    // client.watch().prefix("/session/").create().then(watcher => {
    //     watcher
    //         .on('delete', (res) => {
    //             const session = res.key.toString()
    //             console.log("session delete", session)
    //         })
    //         .on('put', (res) => {
    //             const session = res.key.toString()
    //             const data = res.value.toString()
    //             console.log("session created", session, data)
    //         });
    // });

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
            .on('put', (res) => {
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
