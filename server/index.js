const express = require('express')
const http = require('http');
const cors = require('cors')
const app = express()
const { Etcd3 } = require('etcd3');


console.log("process.env.ETCD", process.env.ETCD)

const client = new Etcd3({
    "hosts": process.env.ETCD
});
const { Server } = require("socket.io");

let avaiable_hosts = {}

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



const PORT = process.env.PORT || 4000
server.listen(PORT, async function () {
    // server ready to accept connections here
    console.log("server has started", PORT)

    const hosts = await client.getAll().prefix('available-hosts/').keys();
    console.log('available hosts:', hosts);

    const session = await client.getAll().prefix('/session/').keys();
    console.log('available sessions:', session);

    client.watch().prefix("/session/").create().then(watcher => {
        watcher
            .on('delete', (res) => {
                const session = res.key.toString()
                console.log("session delete", session)
            })
            .on('put', (res) => {
                const session = res.key.toString()
                const data = res.value.toString()
                console.log("session created", session, data)
            });
    });

    client.watch().prefix("available-hosts/").create().then(watcher => {
        watcher
            // .on('disconnected', () => console.log('disconnected...'))
            // .on('connected', () => console.log('successfully reconnected!'))
            // .on('end', () => {
            //     console.log("ended")
            // })
            // .on("error", () => {
            //     console.log("error")
            // })
            .on('delete', (res) => {
                const host = res.key.toString()
                console.log("delete", host)
                if (avaiable_hosts[host]) {
                    delete avaiable_hosts[host]
                }
            })
            .on('put', (res) => {
                const host = res.key.toString()
                const data = JSON.parse(res.value.toString())
                console.log(host, ' ping ' , data["cpufree"]) //, data
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
