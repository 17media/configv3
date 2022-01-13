# config
Golang process config synchronization using etcd

migrate from 17media/config
and update etcdclient to v3

deprecates 17media/go-etcd

# TODO

for pusher, 
only push modified files as key to etcd, currently we push all keys in repo

for clientv3,
we need to check behavior when etcd is down (should reconnect?)

documentation

check stats implementation