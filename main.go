package main

import (
	"os"
	"truck-unii-app/apps/myconfig/myserver"
	_ "truck-unii-app/docs/api"
	"truck-unii-app/pkg/myconnect"
)

func main() {
	app := myserver.New()
	db := myconnect.NewPostgres(os.Getenv("DB_CONNECT"))
	myserver.Route(app, db)
	myserver.Run(app, "8080")
}
