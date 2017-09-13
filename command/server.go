package command

import (
	"github.com/markbest/minifs/server"
)

var (
	ServerPort     = cmdServer.Flag.Int("port", 1234, "master server port")
	ServerSavePath = cmdServer.Flag.String("dir", "/tmp", "master server data save file path")
)

func init() {
	cmdServer.Run = runServer
}

var cmdServer = &Command{
	UsageLine: "server -port=1234 -dir=/tmp",
	Short:     "start a master server",
	Long:      "start a master server to provide storage spaces",
}

func runServer(cmd *Command, args []string) bool {
	//new master server
	server.NewMasterServer(*ServerSavePath, *ServerPort)
	return true
}
