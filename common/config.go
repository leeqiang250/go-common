package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func LoadConfig(path string, conf *interface{}) {
	file, err := ioutil.ReadFile(path) //"conf/config.yaml"
	if err != nil {
		Error.Println(path, "load err", err)
		os.Exit(0)
	}
	err = yaml.Unmarshal(file, conf)
	if err != nil {
		Error.Println(path, "load err", err)
		os.Exit(0)
	}
}
