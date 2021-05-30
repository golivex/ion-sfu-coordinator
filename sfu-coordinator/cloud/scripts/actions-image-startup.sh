#!/bin/sh

sudo su

cd /home/manish/live_ion_cluster
git config pull.rebase false
GIT_SSH_COMMAND='ssh -i /home/manish/id_rsa -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no' git pull

# git clone https://github.com/cryptagon/ion-cluster.git

cat my_password.txt | docker login --username exceltech --password-stdin

IP=$(curl -H "Metadata-Flavor: Google" http://metadata/computeMetadata/v1/instance/network-interfaces/0/access-configs/0/external-ip)
echo $IP

rm -rf docker-compose-actions-gcp.yml
sed "s/{ip}/$IP/g" docker-compose-actions-gcp-template.yml >> docker-compose-actions-gcp.yml

sudo ulimit -c unlimited
sudo ulimit -SHn 1000000
sudo sysctl -w net.ipv4.tcp_keepalive_time=60
sudo sysctl -w net.ipv4.tcp_timestamps=0
sudo sysctl -w net.ipv4.tcp_tw_reuse=1
#sysctl -w net.ipv4.tcp_tw_recycle=0
sudo sysctl -w net.core.somaxconn=65535
sudo sysctl -w net.ipv4.tcp_max_syn_backlog=65535
sudo sysctl -w net.ipv4.tcp_syncookies=1

sudo docker-compose -f docker-compose-actions-gcp.yml pull
sudo docker-compose -f docker-compose-actions-gcp.yml stop
sudo docker-compose -f docker-compose-actions-gcp.yml up -d 