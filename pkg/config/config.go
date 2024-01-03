package config

import (
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Service struct {
		AppMode    string `yaml:"app_mode"`
		ServerPort int    `yaml:"server_port"`
	}
	MySQL struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		User   string `yaml:"user"`
		Passwd string `yaml:"passwd"`
		DBName string `yaml:"dbname"`
	}
}

var AppConfig Config

func Load() {
	file, err := os.Open("/config/workspace/sources/golang/Pharmacy-POS/config/config.yaml")
	if err != nil {
		panic("打开配置文件失败")
	}

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		panic("读取配置文件内容失败")
	}

	file.Close()
}

func GetDb() (dsn string) {
	dsn = AppConfig.MySQL.User + ":" + AppConfig.MySQL.Passwd + "@tcp("
	dsn += AppConfig.MySQL.Host + ":" + strconv.Itoa(AppConfig.MySQL.Port) + ")/"
	dsn += AppConfig.MySQL.DBName
	return dsn
}
