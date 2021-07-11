# live_ion_cluster

1) The idea here is that, will be using https://github.com/cryptagon/ion-cluster on google cloud instances which will auto scale via cloud instance load balancer

2) Now the problem is that, when we add a new container, we would need to know the ip address of that container/number of sessions it has so that when a new call/conference
is start the frontend is able to connect to the correct node

Solution 
================

This is a crude solutions as of now

1) Sfu has been modified and a hearbeat has been to it. so that it will notify etcd service that its active. it will also 
update its active sessions on it

2) Will have a master http service which will have a list of nodes, based on the load/no of session / existing sessions via etcd it will provide 
the node in which we can start a new session

3.a) we will have to identify the type of call as well i.e group/call (small max 6), meeting upto 50, conference (200-300). if is type of call is group, we can allow to join even if load is more than 80% but not for meeting and conference. but there are lot of issues in this... need to think of a better way

4) we need to know capablity of a server before hand, that is the best way to handle this and also the requirment of a call

if we know its a server than can handle only 20hosts, there is no point assiging it a conference for 300 users
like if we have a meeting room with 30 people, we should start a cpu with 2vcpu and assign it to that host only
if we have small group we can use n1 instances

===  how above will be implemented, as soon we get a request we start a server based on required capacity. suppose we get a request for 50 users we start a server. in the same time we get another user for 50 users on a different session what happens. we will have to wait for the first server to start and second session request will have to wait for almost 40sec and then only its server will start this a problem...............

also when new server will start it will automatically get allocated to any of the sessions if we don't implement any kind of hold on server which i am not planning to... that is also fine.... 

what if i start server for every new session insteadnly. that would be best..... 

but support if i get a request for 100 users and 30 users... not i have started two servers... one for 100 and 30 users.. 
but problem is that what can happen is that 100 users server get allocated to 30users server, and 100 users session is empty.. so it will take more time.

so we will have to block server for session and also we have to start multiple servres together as needed...
both are needed


also suppose we have an n1 server, and lot we get lot of connections. then cpu goes of 100 easily.
we need to way to slow down how we give out hosts. and wait for actual cpu reply to come?


4.a) there another issues suppose we have instance with 16 cpu and one instance just 1 cpu. but 16 cpu cpu load is hight ets say only 10% but that 1 cpu is 1%, our algorightm will choose 1cpu. this is a problem... out assigning host algo should look at cpu load also

4) for live stream in which there is only publisher and multiple subscribers, we can forward sfu from one server to another. new users will have to wait till the new server is active. 

5) in live stream, if we are doing live stream or recording etc we will start a new instance only for this purpose always


How to Scale Horizontal
========================

GCP server will automatically start a new gcp server is more load is required, etcd will ensure that unique session accross sfu

problem, is right now we are dividing session as per lower load, so if 5 gcp server were started sessions will get assigned accross these 5 servers. even if all sessions are very small, it wont happen that first server is used fully and then the second server......... 

need to find a solution for this as well



How To Scale Thoughts Vertical
===============================

Lets say we have sfu A, B, C, D and every sfu can take a load of 2 tracks only

diffent cases that can happen

a) we have lot of incoming requests together on the cluster

A -> 2

then A needs to start mirror from A -> B

need to wait for B to start session and then allocate users to B

but how to allocate users as max load is 2 and have multiple income requests. if we allow more than sfu might crash
.....
this would mean we need to apply locks and can allocate hosts one at a time and not parallely

how do we wait for a new sfu to start, in that time not scale a new server (maybe we just start a new server always but interally don't start two servers together)


b) problem, for a peer to start streaming, it takes few seconds for streams to start. so if we have 100 requests togther load will go in few sections but till then we will allocate existing server to all 100 and it might crash how to solve?

Observation
==============

n1-standard-1 upto 20-30 80% load
n2-highcpu-8  upto 40 only 25% load

video 640x360 bigbuckbunny


livestream on dev server with 500 users
2966745 root      20   0 2178204   1.4g  19232 S 173.7   4.6 128:10.19 sfu       
so 500 * 4 = 2k users can be handled on dev server itself




export GO111MODULE=on


go run cmd/signal/allrpc/main.go -c ./../cfgs/sfuA.toml -jaddr :7002 -gaddr :50052 -eaddr 0.0.0.0:2379 -ipaddr 5.9.18.28


Caddy
=======

Provide the cloudflare token to edit DNS zones and not global api key