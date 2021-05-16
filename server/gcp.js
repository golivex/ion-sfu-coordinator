const { exec } = require("child_process");
var { DateTime } = require('luxon');

const MAX_INSTANCES = process.env.MAX_GCP_INSTANCES || 2

const runCommand = (cmd) => {
    console.log(cmd)
    return new Promise((resolve, reject) => {
        exec(cmd, (error, stdout, stderr) => {
            if (error) {
                console.log(`error: ${error.message}`);
                reject(error)
                return;
            }
            if (stderr) {
                // console.log(`stderr: ${stderr}`);
                resolve(stderr)
                return;
            }
            // console.log(`stdout: ${stdout}`);
            resolve(stdout)
        });
    })
}

export const startServer = async (instance_name = "sfu", zone_idx = 0) => {

    let zone = getZones()[zone_idx]
    let json = await getInstanceList()
    json = JSON.parse(json)
    json = json.filter(inst => inst["name"].indexOf("sfu") !== -1)

    console.log("number of instances running", json.length)
    if (json.length < MAX_INSTANCES) {
        const start_instance_name = instance_name + "-" + new Date().getTime()
        let cmd = `gcloud beta compute instances create ${start_instance_name} --zone=${zone} --tags=sfu --image-family=ubuntu-2004-lts --image-project=ubuntu-os-cloud --maintenance-policy=TERMINATE  --machine-type=n1-standard-2 --boot-disk-type=pd-ssd --metadata-from-file startup-script=/usr/src/app/startup.sh --create-disk size=100GB,type=pd-ssd,auto-delete=yes --format=json` //--scopes=logging-write,compute-rw,cloud-platform
        try {
            const resp = await runCommand(cmd)
            console.log("resp", resp)
            if (resp.indexOf("Created") !== -1) {
                let json = await getInstanceList()
                json = JSON.parse(json)
                const current_instance = json.find(inst => inst["name"] === start_instance_name)
                if (!current_instance) {
                    throw new Error("current created instance not found some major issue")
                }
                console.log("instanced created!!!", current_instance)
                return current_instance
            } else {
                if (zone_idx < getZones().length)
                    return await startServer(instance_name, zone_idx + 1)
                else {
                    console.log("exhuasted all zones")
                    return false
                }
            }
        } catch (error) {
            console.error(error)
            return false
        }

    } else {
        console.log("cannot start more instances")
    }

}

export const deleteServer = (instance_name, zone) => {
    let cmd = `gcloud beta compute instances delete ${instance_name} --zone=${zone} --quiet --format=json` //--zone=${zone}
    return runCommand(cmd)
}

export const getInstanceList = () => {
    let cmd = `gcloud beta compute instances list --format=json ` //  --filter='tags:sfu'
    return runCommand(cmd)
}

export const getZones = () => {
    // let cmd = `gcloud beta compute zones list`
    // return runCommand(cmd)
    return ["asia-south1-a", "asia-south1-b", "asia-south1-c", "asia-east1-a", "asia-east1-b", "asia-east1-c", "us-central1-a", "us-central1-b"]
}