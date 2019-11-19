package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Yaml struct {
	Port string `yaml:"port"`
	Mysql struct {
		User string `yaml:"user"`
		Host string `yaml:"host"`
		Password string `yaml:"password"`
		Port string `yaml:"port"`
		Db string `yaml:"db"`
	}
	YunPian struct {
		ApiKey string `yaml:"apikey"`
	}
	YunXin struct {
		AppKey string `yaml:"app_key"`
		AppSecret string `yaml:"app_secret"`
	}
}

func GetConf () *Yaml {
	yamlFile, err := ioutil.ReadFile("../conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	conf := new(Yaml)
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return conf
}
