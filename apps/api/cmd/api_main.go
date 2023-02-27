package main

import (
	"chat/apps/api/global"
	"chat/apps/api/internal/db"
	"chat/apps/api/internal/router"
)

func main() {
	global.InitConfig()
	db.InitMysql()
	router.InitRouters()
}
