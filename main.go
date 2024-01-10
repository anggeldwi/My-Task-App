package main

import (
	"my-task-api/app/configs"
	"my-task-api/app/databases"
	"my-task-api/app/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := configs.InitConfig()
	dbSql := databases.InitDBMysql(cfg)

	e := echo.New()

	// melakukan migrasi pada database yang telah diinisialisasi
	databases.InitialMigration(dbSql)

	routers.InitRouter(dbSql, e)
	//start server and port
	e.Logger.Fatal(e.Start(":8080"))
}
