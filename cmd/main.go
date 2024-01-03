package main

import (
	"pharmacy-pos/pkg/config"
	"pharmacy-pos/pkg/db"
)

func main() {

	db.InitDB()
	config.Load()

}
