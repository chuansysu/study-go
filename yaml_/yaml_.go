package yaml_

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Config Config
}

type Config struct {
	Models []Model
	Acls   []Acl
	Arrs   []string
}

type Model struct {
	Name   string
	Schema string
}

type Acl struct {
	Model     string
	Role      string
	Operation string
	Action    string
}

func GetConf() Conf {
	var conf Conf
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		//fmt.Println("read file conf.yaml err:", err)
		log.Fatalln("read file conf.yaml err:", err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalln("Unmarshal yaml err:", err)
	}
	return conf
}

func DemoYaml() {
	conf := GetConf()
	log.Println(conf.Config.Models[0].Name)
	log.Println(conf.Config.Models[0].Schema)

	log.Println(conf.Config.Models[1].Name)
	log.Println(conf.Config.Models[1].Schema)

	log.Println(conf.Config.Acls[0].Model)
	log.Println(conf.Config.Acls[0].Role)
	log.Println(conf.Config.Acls[0].Operation)
	log.Println(conf.Config.Acls[0].Action)

	log.Println(conf.Config.Acls[1].Model)
	log.Println(conf.Config.Acls[1].Role)
	log.Println(conf.Config.Acls[1].Operation)
	log.Println(conf.Config.Acls[1].Action)

	log.Println(conf.Config.Arrs)
}
