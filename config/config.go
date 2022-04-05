package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Version = viper.GetString("version")
)

type Project struct {
	ServiceName string
	ServiceAddr string
	ServicePort string
}

func NewProject() *Project {
	return &Project{
		ServiceName: viper.GetString("params.service_name"),
		ServiceAddr: viper.GetString("params.service_addr"),
		ServicePort: viper.GetString("params.service_port"),
	}
}

type Connections struct {
	Host     string
	UserName string
	Password string
	Port     string
	DB       string
	Charset  string
}

func NewConnectins() *Connections {
	return &Connections{
		Host:     viper.GetString("connections.mysql.host"),
		UserName: viper.GetString("connections.mysql.user"),
		Password: viper.GetString("connections.mysql.password"),
		Port:     viper.GetString("connections.mysql.port"),
		DB:       viper.GetString("connections.mysql.db"),
		Charset:  viper.GetString("connections.mysql.charset"),
	}
}

func init() {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("config read error")
	}
}
