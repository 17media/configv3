# config
Golang process config synchronization using etcd

migrate from 17media/config
and update etcdclient to v3

deprecates 17media/go-etcd

note that we ignore errors that files are too large to push (limitation for server side: 1.5 MiB, client side: 4 MiB), because they are not etcd meant to maintain

see
https://etcd.io/docs/v3.3/upgrades/upgrade_3_2/#changed-maximum-request-size-limits-3210 .


# TODO

for pusher, 
only push modified files as key to etcd, currently we push all keys in repo

for clientv3,
we need to check behavior when etcd is down (should reconnect?)

documentation

check stats implementation