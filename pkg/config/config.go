package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Service struct {
		AppMode    string `yaml:"app_mode"`
		ServerPort string `yaml:"server_port"`
	}
	MySQL struct {
		Host   string `yaml:"host"`
		Port   int    `yaml:"port"`
		User   string `yaml:"user"`
		Passwd string `yaml:"passwd"`
		DBName string `yaml:"dbname"`
	}
	Jwt struct {
		Key  string `yaml:"key"`
		Cost int    `yaml:"cost"`
	}
}

var AppConfig Config

// 加载配置文件
func Load() {
	file, err := os.Open("./configs/config.yaml")
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

// 从配置文件获取数据库连接
func GetDb() (dsn string) {
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.MySQL.User,
		AppConfig.MySQL.Passwd,
		AppConfig.MySQL.Host,
		AppConfig.MySQL.Port,
		AppConfig.MySQL.DBName,
	)

	return dsn
}
