package command

import (
	"bytes"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
	. "github.com/markbest/minifs/pb"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	upload     UploadOptions
	serverHost = cmdUpload.Flag.String("host", "127.0.0.1", "master host")
	serverPort = cmdUpload.Flag.Int("port", 1234, "master host port")
)

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
	UsageLine: "upload -host=127.0.0.1 -port=1234 file1 [file2 file3]\n       minifs upload -host=127.0.0.1 -port=1234 -dir=one_directory -include=*.pdf",
	Short:     "upload one or a list of files",
	Long:      "upload one or a list of files",
}

func runUpload(cmd *Command, args []string) bool {
	if len(args) < 1 {
		var files []string
		upload_dir := *upload.dir
		upload_include := *upload.include
		if upload_dir != "" {
			files = getDirFiles(upload_dir, upload_include)
			processUpload(files)
		}
	} else {
		processUpload(args)
	}
	return true
}

//gets all the files under the directory
func getDirFiles(path string, pattern string) []string {
	files := make([]string, 0)
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if !f.IsDir() {
			if pattern != "" {
				if strings.Contains(path, pattern) {
					files = append(files, path)
				}
			} else {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	return files
}

//perform the file upload operation
func processUpload(files []string) {
	for _, file := range files {
		fileHandle, err := os.Open(file)
		if err != nil {
			panic(err)
		}
		file_content, _ := ioutil.ReadAll(fileHandle)

		f := &File{
			Name:    *proto.String(file),
			Content: *proto.String(string(file_content)),
		}

		content, _ := proto.Marshal(f)
		host := *serverHost
		port := strconv.Itoa(*serverPort)
		url := "http://" + host + ":" + port + "/upload"
		req, _ := http.NewRequest("POST", url, bytes.NewReader(content))
		resp, _ := http.DefaultClient.Do(req)

		defer resp.Body.Close()
		_, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("success upload file: " + file)
		}
	}
}
