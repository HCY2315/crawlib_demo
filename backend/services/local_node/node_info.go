package local_node

import (
	"errors"
	"os"

	"github.com/hashicorp/go-sockaddr"
)

var localNode *LocalNode

type IdentifyType string

const (
	Ip       = IdentifyType("ip")
	Mac      = IdentifyType("mac")
	Hostname = IdentifyType("hostname")
)

type local struct {
	Ip           string
	Mac          string
	Hostname     string
	Identify     string
	IdentifyType IdentifyType
}
type LocalNode struct {
	local
	mongo
}

func (l *LocalNode) Ready() error {
	err := localNode.load(true)
	if err != nil {
		return err
	}
	go localNode.watch()
	return nil
}

func NewLocalNode(ip string, identify string, identifyTypeString string) (node *LocalNode, err error) {
	addrs, err := sockaddr.GetPrivateInterfaces()
	// addrs:[192.168.1.14/24 {6 1500 en0 14:7d:da:83:0b:4d up|broadcast|multicast}]
	if ip == "" {
		if err != nil {
			return node, err
		}
		if len(addrs) == 0 {
			return node, errors.New("address not found")
		}
		ipaddr := *sockaddr.ToIPAddr(addrs[0].SockAddr)
		ip = ipaddr.NetIP().String()
	}

	mac := addrs[0].HardwareAddr.String()
	hostname, err := os.Hostname()
	if err != nil {
		return node, err
	}
	local := local{Ip: ip, Mac: mac, Hostname: hostname}
	switch IdentifyType(identifyTypeString) {
	case Ip:
		local.Identify = local.Ip
		local.IdentifyType = Ip
	case Mac:
		local.Identify = local.Mac
		local.IdentifyType = Mac
	case Hostname:
		local.Identify = local.Hostname
		local.IdentifyType = Hostname
	default:
		local.Identify = identify
		local.IdentifyType = IdentifyType(identifyTypeString)
	}
	return &LocalNode{local: local, mongo: mongo{}}, nil
}
