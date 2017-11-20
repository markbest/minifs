package utils

import (
	"errors"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

var (
	Conf              config
	defaultConfigFile = "conf.toml"
)

type config struct {
	//Zookeeper config
	Zookeeper zkp `toml:"zookeeper"`

	//Redis config
	Redis rds `toml:"redis"`
}

type zkp struct {
	Host    []string `toml:"zk_url"`
	Master  string   `toml:"zk_master"`
	Servers string   `toml:"zk_servers"`
}

type rds struct {
	Host     string `toml:"redis_host"`
	Password string `toml:"redis_password"`
	Port     string `toml:"redis_port"`
	Db       int    `toml:"redis_db"`
}

func InitConfig() (err error) {
	configBytes, err := ioutil.ReadFile(defaultConfigFile)
	if err != nil {
		return errors.New("config load err:" + err.Error())
	}
	_, err = toml.Decode(string(configBytes), &Conf)
	if err != nil {
		return errors.New("config decode err:" + err.Error())
	}
	return nil
}
