package main

import (
	"fmt"
	"pharmacy-pos/pkg/config"
	"pharmacy-pos/pkg/logger"
)

func main() {
	logger.Init()
	log := logger.GetLogger()
	log.Error("这是一条错误消息")
	log.Warn("这是一条警告消息")
	config.Load()

	fmt.Println("AppMode: ", config.AppConfig.Service.AppMode)
	fmt.Println("ServerPort: ", config.AppConfig.Service.ServerPort)

	fmt.Println("Host: ", config.AppConfig.MySQL.Host)
	fmt.Println("Port: ", config.AppConfig.MySQL.Port)
	fmt.Println("User: ", config.AppConfig.MySQL.User)
	fmt.Println("Passwd: ", config.AppConfig.MySQL.Passwd)
	fmt.Println("DBName: ", config.AppConfig.MySQL.DBName)
}
