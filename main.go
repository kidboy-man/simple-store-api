package main

import (
	database "simple-store-api/databases"
	_ "simple-store-api/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	database.InitDB()
	beego.BConfig.RunMode = "dev"
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
