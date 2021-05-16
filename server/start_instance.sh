#!/bin/bash

# not used

gcloud auth activate-service-account --key-file=/workspace/steady-datum-291915-386b3c696fb0.json >&1
gcloud config set project java-ref >&1

echo "$1" >&1
echo "x$2x" >&1

gcloud beta compute instances create $1 --zone=$2 --image-family=ubuntu-2004-lts --image-project=ubuntu-os-cloud --maintenance-policy=TERMINATE --machine-type=n1-standard-2 --boot-disk-type=pd-ssd --metadata-from-file startup-script=/workspace/startup.sh --scopes=logging-write,compute-rw,cloud-platform --create-disk size=100GB,type=pd-ssd,auto-delete=yes --format=json >&1