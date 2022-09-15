package config

import (
	"gopkg.in/go-ini/ini.v1"
	"log"
	"todo_app/utils"
)

// config.iniを読み込むためのstruct
type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	LogFile   string
	Static    string
}

// グローバルで呼び出せるようにする
var Config ConfigList

// mainの前に呼び出す
func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	// config.iniを読み込む
	cfg, err := ini.Load("config.ini")
	// エラーハンドリング
	if err != nil {
		log.Fatalln("err")
	}
	// グローバルに設定したConfigに値を設定する
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"), //初期値設定
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
		Static:    cfg.Section("web").Key("static").String(),
	}
}
