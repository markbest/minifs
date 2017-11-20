package rpc

import (
	"golang.org/x/net/context"
	"log"
	pb "github.com/markbest/minifs/pbs"
	. "github.com/markbest/minifs/storage"
)

type FileRpcServer struct {
	LogHandle *log.Logger
	UploadDir string
}

func (f *FileRpcServer) DoFileUpload(ctx context.Context, in *pb.FileUploadRequest) (*pb.FileUploadResponse, error) {
	r := &pb.FileUploadResponse{}
	server := NewMiniFs(Options{BasePath: f.UploadDir})
	err := server.Write(in.Name, in.Content)
	if err != nil {
		r.Status = false
		r.Message = err.Error()
		return r, err
	} else {
		r.Status = true
		r.Message = ""
		r.File = in.Name
		return r, nil
	}
}
