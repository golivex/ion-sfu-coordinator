# live_ion_cluster

1) The idea here is that, will be using https://github.com/cryptagon/ion-cluster on google cloud instances which will auto scale via cloud instance load balancer

2) Now the problem is that, when we add a new container, we would need to know the ip address of that container/number of sessions it has so that when a new call/conference
is start the frontend is able to connect to the correct node

Solution 
================

This is a crude solutions as of now

1) Will have a slave http service which will do  heartbeat, publish ip, publish load/no of sessions of a node

2) Will have a master http service which will have a list of nodes, based on the load/no of session / existing sessions via etcd it will provide 
the node in which we can start a new session

this way we will be able to have large number of smaller conference with upto 200-300 peers per server (depending on server)
