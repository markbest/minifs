package command

import (
	"log"
	. "github.com/markbest/minifs/server"
	. "github.com/markbest/minifs/utils"
	"strconv"
)

var (
	ServerHost     = cmdNode.Flag.String("host", "127.0.0.1", "server host")
	ServerPort     = cmdNode.Flag.Int("port", 1235, "server port")
	ServerSavePath = cmdNode.Flag.String("dir", "./data", "server data save file path")
	NodeLogHandle  *log.Logger
)

func init() {
	cmdNode.Run = runServer
}

var cmdNode = &Command{
	UsageLine: "node -host=127.0.0.1 -port=1235 -dir=./data",
	Short:     "start a node server",
	Long:      "start a node server to provide storage spaces",
}

func runServer(cmd *Command, args []string) bool {
	//init log handle
	NodeLogHandle := GetLogHandle("node", *ServerHost, *ServerPort)

	//create zk server node
	nodeHost := *ServerHost + ":" + strconv.Itoa(*ServerPort)
	err := ZK.Create(Conf.Zookeeper.Servers+"/"+nodeHost, nil, 1)
	if err != nil {
		NodeLogHandle.Println(err.Error())
	}

	//start node server
	NewNodeServer(*ServerPort, NodeLogHandle, *ServerSavePath)
	return true
}
