package server

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	. "github.com/markbest/minifs/pb"
	. "github.com/markbest/minifs/storage"
	"net/http"
	"strconv"
)

var (
	Server MasterServer
)

type MasterServer struct {
	//File FileServer
	File FileServer

	//route server
	Route RouteServer
}

type FileServer struct {
	ServerPath string
}

type RouteServer struct {
	ServerPort int
}

func NewMasterServer(path string, port int) {
	Server = MasterServer{
		File: FileServer{
			ServerPath: path,
		},
		Route: RouteServer{
			ServerPort: port,
		},
	}

	router := httprouter.New()
	router.POST("/upload", UploadFile)
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(Server.Route.ServerPort), router))
}

func UploadFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	f := &File{}
	body, _ := ioutil.ReadAll(r.Body)
	proto.Unmarshal(body, f)
	server := NewMiniFs(Options{
		BasePath: Server.File.ServerPath,
	})
	server.Write(f.Name, f.Content)
	fmt.Println("upload file: " + f.Name + " complete")
}
