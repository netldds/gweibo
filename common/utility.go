package common

import (
	"io/ioutil"
	"log"

	"github.com/go-yaml/yaml"
)

var Config *ServiceConfig

func LoadConf() (*ServiceConfig, error) {
	b, err := ioutil.ReadFile("./gweibo.yaml")
	if err != nil {
		log.Fatalln(err)
		return nil, nil
	}
	err = yaml.Unmarshal(b, &Config)
	if err != nil {
		log.Fatal(err)
		return nil, nil
	}
	return Config, nil
}
func ReloadConfig() {
	if s, err := LoadConf(); err == nil {
		Config = s
	}
}
