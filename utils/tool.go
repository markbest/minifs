package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//gets all the files under the directory
func GetDirFiles(path string, pattern string) []string {
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

//get lod handle
func GetLogHandle(log_type string, host string, port int) *log.Logger {
	LogHandle := log.New(os.Stdout, "["+log_type+"-"+host+":"+strconv.Itoa(port)+"]", log.LstdFlags)
	return LogHandle
}
