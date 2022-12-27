package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type conf struct {
	Host string `yml:"host"`
	Port uint   `yml:"port"`
}

func GetConf() conf {
	file, err := ioutil.ReadFile("conf.yml")
	if err != nil {
		log.Fatal(err)
	}
	data := conf{}
	error := yaml.Unmarshal([]byte(file), &data)
	if error != nil {
		log.Fatal(err)
	}
	return data
}
