#!/bin/sh

sudo su

mkdir -p /home/manish

cd /home/manish
sudo apt update
sudo apt upgrade -y
sudo apt install apt-transport-https ca-certificates curl software-properties-common -y
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal stable" -y
sudo apt update
sudo apt-get install curl git docker-ce -y ufw
sudo docker --version

sudo curl -L "https://github.com/docker/compose/releases/download/1.27.4/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose


echo "-----BEGIN OPENSSH PRIVATE KEY-----
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
-----END OPENSSH PRIVATE KEY-----" >> /home/manish/id_rsa

echo "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQC/WNrLuL1D5x8ZcNtrmKVdAURY7S3paL6GIgRe3UKyYNWEW8EqsepRlbNlO6WOAvnHMu6OyydioPe5hzUdyVwB4hiywO2SClGlooIxauFNhcKUNUHR9PIgMxQi/YEbZtiaQFDxrI3ek6Pn/+QZxusDtcY87Wh+QRHwBA4yAqFHciRvNKoh91ehPdbyuxpIk03Z1f4NR3o9J+RFR/iLG4GSsoTJQaTFAS76UNPuDyi0DH/JGhvOD+L+uYgCE7ASI4qLvW3Gm+hyJ+Ea9NZoVdJzW+rogaEWEwEVw9QB7qYuNAvXhZbvd/zU1f0hW00tHSG6KUZRL1wa+HGBZA6Ou8XqyvxMYgFM4PX7wUWnVZo9ARh0OjedsAtVUx9RuTxX5s+lwyWBdwiyWBclpNVBYtd7A8G+aVDQD7Dtcf2f67FYUSXFN0f5wMRu+aFnpfsB6LWq53R4V719Vv5l9+d8vzxDgpZgrxAzwNMQE9w0JHJwwsRWepWmwCSQrm/2GrGItFnJ2ED13DJ0+WDVJ044EJ0lE/rE6K8Ty7m2Rc2oMxLSx6cepyzEWqOctjnYkwv9v8ZX1MhYxc5sQdK9PJwR2XFrkm9R3i3B1/liXY5IkPhni9yiyl/qSY0Z6+Atopc6x63p4YiDskTmxyo+2YdKlNBrgIhdL9R7WcpcabFY/AOVbQ== manish@ryzne-server.com" >> /home/manish/id_rsa.pub

sudo chmod 600 /home/manish/id_rsa
sudo chmod 600 /home/manish/id_rsa.pub

cd /home/manish

GIT_SSH_COMMAND='ssh -i /home/manish/id_rsa -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no' git clone git@github.com:manishiitg/live_ion_cluster.git

sudo ufw allow 22
sudo ufw allow 80
sudo ufw allow 7001
sudo ufw allow 50051
sudo ufw allow 10000:15000/udp
sudo ufw enable

cd /home/manish/live_ion_cluster

# git clone https://github.com/cryptagon/ion-cluster.git

cat my_password.txt | docker login --username exceltech --password-stdin

IP=$(curl -H "Metadata-Flavor: Google" http://metadata/computeMetadata/v1/instance/network-interfaces/0/access-configs/0/external-ip)
echo $IP

sed "s/{ip}/$IP/g" docker-compose-gcp-template.yml >> docker-compose-gcp.yml
sed "s/{ip}/$IP/g" ./cfgs/sfu-gcp.template.toml >> ./cfgs/sfu-gcp.toml

sudo ulimit -c unlimited
sudo ulimit -SHn 1000000
sudo sysctl -w net.ipv4.tcp_keepalive_time=60
sudo sysctl -w net.ipv4.tcp_timestamps=0
sudo sysctl -w net.ipv4.tcp_tw_reuse=1
#sysctl -w net.ipv4.tcp_tw_recycle=0
sudo sysctl -w net.core.somaxconn=65535
sudo sysctl -w net.ipv4.tcp_max_syn_backlog=65535
sudo sysctl -w net.ipv4.tcp_syncookies=1

# sudo docker-compose -f docker-compose-gcp.yml build
sudo docker-compose -f docker-compose-gcp.yml up -d 