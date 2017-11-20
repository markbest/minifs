package command

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io/ioutil"
	pb "github.com/markbest/minifs/pbs"
	. "github.com/markbest/minifs/utils"
	"os"
)

var upload UploadOptions

type UploadOptions struct {
	dir     *string
	include *string
}

func init() {
	cmdUpload.Run = runUpload
	upload.dir = cmdUpload.Flag.String("dir", "", "Upload the whole folder recursively if specified.")
	upload.include = cmdUpload.Flag.String("include", "", "pattens of files to upload, e.g., *.pdf, *.html, works together with -dir")
}

var cmdUpload = &Command{
	UsageLine: "upload file1 [file2 file3]\n       minifs upload -dir=one_directory -include=*.pdf",
	Short:     "upload one or a list of files",
	Long:      "upload one or a list of files",
}

func runUpload(cmd *Command, args []string) bool {
	if len(args) < 1 {
		var files []string
		if *upload.dir != "" {
			files = GetDirFiles(*upload.dir, *upload.include)
			processUpload(files)
		}
	} else {
		processUpload(args)
	}
	return true
}

//perform the file upload operation
func processUpload(files []string) {
	master, _, err := ZK.Get(Conf.Zookeeper.Master)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		handle, err := os.Open(file)
		if err != nil {
			panic(err)
		}

		content, _ := ioutil.ReadAll(handle)
		d := &pb.DispatchRequest{"upload", &pb.FileUploadRequest{*proto.String(file), *proto.String(string(content))}}
		conn, err := grpc.Dial(string(master), grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		c := pb.NewDispatchClient(conn)
		r, err := c.DispatchJobs(context.Background(), d)
		if err != nil {
			panic(err)
		}
		fmt.Println(r)
	}
}
