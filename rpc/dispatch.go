package rpc

import (
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"math/rand"
	pb "github.com/markbest/minifs/pbs"
	. "github.com/markbest/minifs/utils"
	"time"
	"log"
)

type DispatchRpcServer struct{
	LogHandle *log.Logger
}

func (d *DispatchRpcServer) DispatchJobs(ctx context.Context, in *pb.DispatchRequest) (*pb.FileUploadResponse, error) {
	r := &pb.FileUploadResponse{}
	switch in.Action {
	case "upload":
		servers, err := ZK.Children(Conf.Zookeeper.Servers)
		if err != nil {
			return r, err
		}

		if len(servers) > 0 {
			rd := rand.New(rand.NewSource(time.Now().UnixNano()))
			nodeServer := string(servers[rd.Intn(len(servers))])
			conn, err := grpc.Dial(nodeServer, grpc.WithInsecure())
			if err != nil {
				return r, err
			}
			defer conn.Close()

			d.LogHandle.Println("Dispatch upload job to", nodeServer)
			c := pb.NewFileUploadClient(conn)
			r, err := c.DoFileUpload(context.Background(), in.Content)
			if err != nil {
				return r, err
			}
			d.LogHandle.Println("Process job response", r)
			return r, nil
		} else {
			return r, errors.New("No available node server")
		}
	default:
		return r, errors.New("Illegimate action")
	}
}
