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

3) (edit, need to test mirror sfu) we cannot scale up a conference call in which there are multiple publishers. one conf call cannot scale more than a single server, so if server goes to 100% load we simply cannot allow more users we can only allow subscribers after like 80% load.

3.a) we will have to identify the type of call as well i.e group/call (small max 6), meeting upto 50, conference (200-300). if is type of call is group, we can allow to join even if load is more than 80% but not for meeting and conference. but there are lot of issues in this... need to think of a better way

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

n1-standard-2  upto 20-30 clients cpu goes 90%
n1-standard-4  upto 50 clients cpu goes 80%
n2-standard-4  upto 50 clients cpu goes 70%

video 640x360 bigbuckbunny


livestream on dev server with 500 users
2966745 root      20   0 2178204   1.4g  19232 S 173.7   4.6 128:10.19 sfu       
so 500 * 4 = 2k users can be handled on dev server itself


Testing Videos
================
https://test-videos.co.uk/bigbuckbunny/webm-vp9


export GO111MODULE=on


go run cmd/signal/allrpc/main.go -c ./../cfgs/sfuA.toml -jaddr :7002 -gaddr :50052 -eaddr 0.0.0.0:2379 -ipaddr 5.9.18.28