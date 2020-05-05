package common

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

var Config *ServiceConfig

func LoadConf() (s *ServiceConfig, err error) {
	b, err := ioutil.ReadFile("./gweibo.yaml")
	if err != nil {
		log.Fatalln(err)
		return
	}
	err = yaml.Unmarshal(b, &s)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
func ReloadConfig() {
	if s, err := LoadConf(); err == nil {
		Config = s
	}
}
