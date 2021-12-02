package local_node

import (
	"crawlab/model"

	"github.com/spf13/viper"
)

func GetLocalNode() *LocalNode {
	return localNode
}
func CurrentNode() *model.Node {
	return GetLocalNode().Current()
}

func InitLocalNode() (node *LocalNode, err error) {
	// register:
	// # type 填 mac/ip/customName, 如果是ip，则需要手动指定IP, 如果是 customName, 需填写你的 customNodeName
	// type: "mac"
	// customNodeName: "" # 自定义节点名称, default node1,只有在type = customName 时生效
	// ip: ""
	ip := viper.GetString("server.register.ip")
	customNodeName := viper.GetString("server.register.customNodeName")
	registerType := viper.GetString("server.register.type")

	localNode, err = NewLocalNode(ip, customNodeName, registerType)
	if err != nil {
		return nil, err
	}
	return localNode, err
}
