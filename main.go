package main

import (
	"truck-unii-app/apps/myconfig/myserver"
	"truck-unii-app/apps/myconfig/myvar"
	_ "truck-unii-app/docs/api"
	"truck-unii-app/pkg/myconnect"
)

func main() {
	myvar.SetEnv()
	app := myserver.New()
	db := myconnect.NewPostgres(myvar.ENV_DB_CONNECT)
	myserver.InitDB(db)
	myserver.Route(app, db)
	myserver.Run(app, myvar.ENV_PORT)
}
