package config

import (
	"log"
	"main/utils"

	"gopkg.in/go-ini/ini.v1"
)


type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
}

func init(){
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

var Config ConfigList

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln(err)
	}
	Config = ConfigList{
		Port: 			cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: 	cfg.Section("db").Key("driver").String(),
		DbName: 		cfg.Section("db").Key("name").String(),
		LogFile: 		cfg.Section("web").Key("logfile").String(),
	}
}