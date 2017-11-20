package server

import (
	"google.golang.org/grpc"
	"log"
	pb "github.com/markbest/minifs/pbs"
	"github.com/markbest/minifs/rpc"
	"net"
	"strconv"
)

func NewNodeServer(port int, logHandle *log.Logger, uploadDir string) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		logHandle.Println(err.Error())
	}
	logHandle.Println("Start node node complete")
	s := grpc.NewServer()
	pb.RegisterFileUploadServer(s, &rpc.FileRpcServer{LogHandle: logHandle, UploadDir: uploadDir})
	s.Serve(lis)
}
