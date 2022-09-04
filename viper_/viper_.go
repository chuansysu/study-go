package viper_

import (
	"log"

	"github.com/spf13/viper"
)

func GetViper() *viper.Viper {
	return viper.GetViper()
}

func DemoViper() {
	//viper.SetConfigFile("./config.yaml")
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("read config err:", err)
	}
	//log.Println(viper.AllKeys())
	//viper.WriteConfigAs("./myConfig")
}
