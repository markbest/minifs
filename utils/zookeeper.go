package utils

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"os"
	"strings"
	"time"
)

var ZK *Zookeeper

type Zookeeper struct {
	client *zk.Conn
}

type ZookeeperLog struct{}

func (l ZookeeperLog) Printf(str string, params ...interface{}) {
	LogHandle := log.New(os.Stdout, "[zk]", log.LstdFlags)
	LogHandle.Println(str, params)
}

//New ZK connect
func NewZKConn(servers []string) *Zookeeper {
	c, _, err := zk.Connect(servers, time.Second*10, func(co *zk.Conn) {
		co.SetLogger(ZookeeperLog{})
	})
	if err != nil {
		panic(err)
	}
	ZK = &Zookeeper{client: c}
	return ZK
}

//ZK Create new node
//flag：0:forever;1:ephemeral
func (z *Zookeeper) Create(path string, data []byte, flag int32) (err error) {
	var s string = ""
	path_arr := strings.Split(path, "/")
	if len(path_arr) > 1 {
		for _, v := range path_arr {
			if v != "" {
				s = s + "/" + v
				exist, _ := z.Exist(s)
				if !exist {
					if path != s {
						_, err = z.client.Create(s, nil, 0, zk.WorldACL(zk.PermAll))
						if err != nil {
							return err
						}
					} else {
						_, err = z.client.Create(s, data, flag, zk.WorldACL(zk.PermAll))
						if err != nil {
							return err
						}
					}
				}
			}
		}
	} else {
		_, err = z.client.Create(path, data, flag, zk.WorldACL(zk.PermAll))
		if err != nil {
			return err
		}
	}
	return nil
}

//ZK Update node
func (z *Zookeeper) Update(path string, data []byte, version int32) (err error) {
	_, err = z.client.Set(path, data, version)
	if err != nil {
		return err
	}
	return nil
}

//ZK Delete node
func (z *Zookeeper) Delete(path string, version int32) (err error) {
	err = z.client.Delete(path, version)
	if err != nil {
		return err
	}
	return nil
}

//ZK Get node
func (z *Zookeeper) Get(path string) (s []byte, stat *zk.Stat, err error) {
	s, stat, err = z.client.Get(path)
	if err != nil {
		return s, stat, err
	}
	return s, stat, nil
}

//ZK Exist node
func (z *Zookeeper) Exist(path string) (flag bool, err error) {
	flag, _, err = z.client.Exists(path)
	if err != nil {
		return false, err
	}
	return flag, nil
}

//ZK Exist watch event
func (z *Zookeeper) ExistW(path string) (flag bool, ch <-chan zk.Event, err error) {
	flag, _, ch, err = z.client.ExistsW(path)
	return flag, ch, err
}

//ZK Get children
func (z *Zookeeper) Children(path string) (child []string, err error) {
	child, _, err = z.client.Children(path)
	if err != nil {
		return child, err
	}
	return child, nil
}

//ZK Get children watch event
func (z *Zookeeper) ChildrenW(path string) (child []string, ch <-chan zk.Event, err error) {
	child, _, ch, err = z.client.ChildrenW(path)
	return child, ch, err
}

//ZK Close
func (z *Zookeeper) Close() {
	z.client.Close()
}
