package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"mudiralmaham/models"
	"mudiralmaham/router"
	logger "mudiralmaham/utils/Logger"
	"mudiralmaham/utils/authentication"
	"mudiralmaham/utils/database"
	"mudiralmaham/utils/worker"
	"net/http"
	"os"
	"time"
)

func main() {
	router.ApiMapper()
	database.DatabaseInit()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
		logger.GeneralLogger.Printf("Defaulting to port %s\n", port)
	}
	authentication.NewTask = make(chan models.Task)
	go worker.SendNotification(authentication.NewTask)
	srv := &http.Server{
		//Handler: handlers.LoggingHandler(logger.GeneralLogger.Writer(), router.Router),
		Handler:      handlers.LoggingHandler(os.Stdout, router.Router),
		Addr:         "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	startUpLog(false, srv, port)

}

func startUpLog(inFile bool, srv *http.Server, port string) {
	if inFile {
		logger.GeneralLogger.Printf("Listening on port %s\n", port)
		logger.GeneralLogger.Printf("Open http://localhost:%s in the browser\n", port)
		logger.ErrorLogger.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	} else {
		log.Printf("Listening on port %s\n", port)
		log.Printf("Open http://localhost:%s in the browser\n", port)
		log.Fatal(srv.ListenAndServe())
	}
}

func cleanUp() {
	database.DisconnectDB()
}
