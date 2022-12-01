package main

import (
	"gotrading/app/controllers"
	"gotrading/config"
	"gotrading/utils"
)

func main() {
	utils.LoggingSettings((config.Config.LogFile))
	// apiCient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	// fmt.Println(models.DbConnection)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
	// fmt.Println(apiCient.GetBalance())
	// fmt.Println(models.DbConnection)
}
