package server

import (
	"google.golang.org/grpc"
	"log"
	pb "github.com/markbest/minifs/pbs"
	"github.com/markbest/minifs/rpc"
	"net"
	"strconv"
)

func NewMasterServer(port int, logHandle *log.Logger) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		logHandle.Println(err.Error())
	}
	logHandle.Println("Start master node complete")
	s := grpc.NewServer()
	pb.RegisterDispatchServer(s, &rpc.DispatchRpcServer{LogHandle:logHandle})
	s.Serve(lis)
}
