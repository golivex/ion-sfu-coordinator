const { exec } = require("child_process");


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
                console.log(`stderr: ${stderr}`);
                try {
                    resolve(JSON.parse(stderr))
                } catch (error) {
                    throw error
                }
                return;
            }
            console.log(`stdout: ${stdout}`);
            try {
                resolve(JSON.parse(stdout))
            } catch (error) {
                throw error
            }

        });
    })
}

export const startServer = async (instance_name, zone = false) => {
    if (!zone) {
        zone = getZones()[0]
    }
    let cmd = `gcloud beta compute instances create ${instance_name} --zone=${zone} --tags=sfu --image-family=ubuntu-2004-lts --image-project=ubuntu-os-cloud --maintenance-policy=TERMINATE  --machine-type=n1-standard-2 --boot-disk-type=pd-ssd --metadata-from-file startup-script=/usr/src/app/startup.sh --create-disk size=100GB,type=pd-ssd,auto-delete=yes --format=json` //--scopes=logging-write,compute-rw,cloud-platform
    return runCommand(cmd)

}

export const deleteServer = (instance_name, zone) => {
    let cmd = `gcloud beta compute instances delete ${instance_name} --quiet --format=json` //--zone=${zone}
    return runCommand(cmd)
}

export const getInstanceList = () => {
    let cmd = `gcloud beta compute instances list --filter='tags:sfu' --format=json `
    return runCommand(cmd)
}

export const getZones = () => {
    // let cmd = `gcloud beta compute zones list`
    // return runCommand(cmd)
    return ["asia-south1-a", "asia-south1-b", "asia-south1-c", "asia-east1-a", "asia-east1-b", "asia-east1-c", "us-central1-a", "us-central1-b"]
}