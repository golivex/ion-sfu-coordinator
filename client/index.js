const express = require('express')
const http = require('http');
const cors = require('cors')
const { createTerminus } = require('@godaddy/terminus');
const app = express()
const osu = require('node-os-utils')
const { Etcd3 } = require('etcd3');
const client = new Etcd3({
    "hosts" : process.env.ETCD
});

const HEALTH_INTERVAL = process.env.HEALTH_INTERVAL || 5000
const lease = client.lease(HEALTH_INTERVAL * 2 / 1000)

app.use(cors())

app.get('/', (req, res) => {
    res.send('ping client')
})


function onSignal() {
    console.log('server is starting cleanup');
    return Promise.all([
        // your clean logic, like closing database connections
    ]);
}


async function onShutdown() {
    console.log('cleanup finished, server is shutting down');
    await updateShutdown()

}
const cpu = osu.cpu
const mem = osu.mem
const os = osu.os

let instanceMetadata = {}
let gcp_ip = {}
const ip = process.env.IP || os.ip()

console.log("ip", ip)

function healthCheck() {

    return new Promise(async (resolve) => {

        const cpufree = await cpu.free()
        const cpuusage = await cpu.usage()
        const memfree = await mem.info()
        resolve({
            "count": cpu.count(),
            "cpuusage": cpuusage,
            "cpufree": cpufree,
            "model": cpu.model(),
            "memory": memfree,
            "hostname": os.hostname(),
            "instanceMetadata": instanceMetadata,
            "identify": getIdentifyData()
        })


    })
}
const server = http.createServer(app);

const beforeShutdown = () => {
    console.log("beforeShutdown")
}

const options = {
    // health check options
    healthChecks: {
        '/healthcheck': healthCheck,    // a function returning a promise indicating service health,
        verbatim: true, // [optional = false] use object returned from /healthcheck verbatim in response,
        __unsafeExposeStackTraces: true // [optional = false] return stack traces in error response if healthchecks throw errors
    },

    // cleanup options
    beforeShutdown,
    timeout: 1000,                   // [optional = 1000] number of milliseconds before forceful exiting
    onSignal,                        // [optional] cleanup function, returning a promise (used to be onSigterm)
    onShutdown,                      // [optional] called right before exiting
    // both
};

createTerminus(server, options);

const PORT = process.env.PORT || 4001

const getKey = () => {
    const data = getIdentifyData()
    const key = 'available-hosts/' + data.ip + ":" + data.port
    return key
}

const getIP = () => {
    return ip
}

const getIdentifyData = () => {
    return {
        ip: getIP(),
        port: process.env.SFU_PORT
    }
}

const updateIdentify = async () => {
    const data = await healthCheck()
    await lease.keepaliveOnce()
    await lease.put(getKey()).value(JSON.stringify(data));
}

const updateShutdown = async () => {
    await lease.revoke()
    await client.delete().key(getKey())
}


server.listen(PORT, async function () {
    // server ready to accept connections here
    console.log("client has started", PORT)
    await updateIdentify()

    setInterval(async () => {
        // console.log("updating health check")
        await updateIdentify()
    }, HEALTH_INTERVAL)

});
