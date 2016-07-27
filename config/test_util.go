package config

const (
	testConfig = `
bootstrapper: 10.10.10.192

subnet: 10.10.10.0
netmask: 255.255.255.0
iplow: 10.10.10.100
iphigh: 10.10.10.199
routers: [10.10.10.192]
broadcast: 10.10.10.255
nameservers: [10.10.10.192, 8.8.8.8, 8.8.4.4]
domainname: unisound.com

nodes:
  - mac: "00:25:90:c0:f7:80"
    ip: 10.10.10.201
    ceph_monitor: y
    kube_master: y
    etcd_member: y
  - mac: "00:25:90:c0:f6:ee"
    ip: 10.10.10.202
    ceph_monitor: y
    etcd_member: y
  - mac: "00:25:90:c0:f6:d6"
    ceph_monitor: y
    etcd_member: y
  - mac: "00:25:90:c0:f7:ac"
    ip: "10.10.10.204"
  - mac: "00:25:90:c0:f7:7e"
    ip: "10.10.10.205"
`
)
