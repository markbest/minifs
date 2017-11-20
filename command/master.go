package command

import (
	"github.com/samuel/go-zookeeper/zk"
	"log"
	. "github.com/markbest/minifs/server"
	. "github.com/markbest/minifs/utils"
	"os"
	"strconv"
	"time"
)

var (
	MasterHost      = cmdMaster.Flag.String("host", "127.0.0.1", "master server host")
	MasterPort      = cmdMaster.Flag.Int("port", 1234, "master server port")
	MasterLogHandle *log.Logger
)

func init() {
	cmdMaster.Run = runMaster
}

var cmdMaster = &Command{
	UsageLine: "master -host=127.0.0.1 -port=1234",
	Short:     "start a master server",
	Long:      "start a master server to provide storage services",
}

func runMaster(cmd *Command, args []string) bool {
	//init log handle
	MasterLogHandle = GetLogHandle("master", *MasterHost, *MasterPort)

	flag, err := ZK.Exist(Conf.Zookeeper.Master)
	if err != nil {
		MasterLogHandle.Println(err.Error())
		os.Exit(1)
	}

	if flag {
		tick := time.NewTicker(time.Second * 10)
		for {
			select {
			case currentTime := <-tick.C:
				_, ch, err := ZK.ExistW(Conf.Zookeeper.Master)
				if err != nil {
					MasterLogHandle.Println(err.Error())
					os.Exit(1)
				}

				event := <-ch
				if event.Type == zk.EventNodeDeleted {
					MasterLogHandle.Println("Detecte master node is closed, retry create master node:", *MasterHost, *MasterPort)
					err = createMasterAndStartMaster()
					if err != nil {
						MasterLogHandle.Println(err.Error(), currentTime.Format("2006-01-02 15:04:05"))
					} else {
						NewMasterServer(*MasterPort, MasterLogHandle)
					}
				}
			}
		}
	} else {
		MasterLogHandle.Println("Detecte master node not exist, start create master node:", *MasterHost, *MasterPort)
		err = createMasterAndStartMaster()
		if err != nil {
			MasterLogHandle.Println(err.Error())
		} else {
			NewMasterServer(*MasterPort, MasterLogHandle)
		}
	}
	return true
}

func createMasterAndStartMaster() error {
	masterHost := *MasterHost + ":" + strconv.Itoa(*MasterPort)
	err := ZK.Create(Conf.Zookeeper.Master, []byte(masterHost), 1)
	if err != nil {
		return err
	}
	return nil
}
