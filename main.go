package main

import (
	"net/http"

	"github.com/bino1490/case-study/api/handler"
	"github.com/bino1490/case-study/pkg/config"
	"github.com/bino1490/case-study/pkg/entity"
	"github.com/bino1490/case-study/pkg/logger"
	"github.com/bino1490/case-study/pkg/repository"
	"github.com/bino1490/case-study/pkg/service"
)

//-- Main to Inatialize the service
func main() {
	logger.BootstrapLogger.Info("Service starting...")
	initService()
}

//-- initService initialize the service ----
func initService() {
	logger.BootstrapLogger.Debug("Entering initService...")
	logger.BootstrapLogger.Info("Starting " + config.SrvConfig.GetString("application.name") +
		" with profile=" + config.SrvConfig.GetString("profile") + " properties")

	//repository := initDatabase()
	//service := service.NewService(repository)

	inMemSvc := service.NewInMemService()
	addInMemHandlers(inMemSvc)
	err := http.ListenAndServe(":"+config.SrvConfig.GetString("http.port"), nil)
	if err != nil {
		logger.BootstrapLogger.Error("Failed to ListenPort 8080")
		panic(err)
	}
}

func addInMemHandlers(inMemSvc *service.MemHandlers) {
	//http.HandleFunc("/in-memory", inMemSvc.InMemGetPOST)
	http.Handle("/in-memory", handler.InMemReqHandler(inMemSvc))
	http.Handle("/in-memory/", handler.InMemReqHandler(inMemSvc))
}

//-- To initialize the database ----
func initDatabase() repository.DbRepository {
	logger.BootstrapLogger.Debug("Entering initDatabase...")

	if config.SrvConfig.GetString(
		"database.mongodb.enable") == "true" {
		logger.BootstrapLogger.Debug("About to initialize MongoDB repo...")
		return repository.NewCbRepository()

	} else {
		// Throw panic
		logger.BootstrapLogger.Error("Incorrect database configuration settings. Can't proceed!")
		panic(entity.ErrInvalidConfig)
	}
}

// func addHandlers(mux *http.ServeMux, service service.ScheduleService) {

// 	r.Handle("/getschedule/{channelId}",
// 		middleware.AccessLog(
// 			middleware.ParseHeader(
// 				handler.GetScheduleByChannelID(
// 					service)))).Methods(http.MethodGet, http.MethodOptions)
// }
