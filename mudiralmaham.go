package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mudiralmaham/api"
	"mudiralmaham/utils/Database"
	logger "mudiralmaham/utils/Logger"
	"net/http"
	"os"
	"time"
)

var (
	Router  = mux.NewRouter()
)

func main() {
	apiMapper()
	Database.DatabaseInit()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
		logger.GeneralLogger.Printf("Defaulting to port %s\n", port)
	}

	srv := &http.Server{
		Handler: Router,
		Addr: "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	startUpLog(false, srv, port)
}
/**
	#main api router
		it maps incoming routes to dedicated functions
 */
func apiMapper() {
	auth := Router.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", api.Login)
	auth.HandleFunc("/signUp", api.SignUp)
}

func startUpLog(inFile bool, srv *http.Server, port string) {
	if inFile {
		logger.GeneralLogger.Printf("Listening on port %s\n", port)
		logger.GeneralLogger.Printf("Open http://localhost:%s in the browser\n", port)
		logger.ErrorLogger.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	} else{
		log.Printf("Listening on port %s\n", port)
		log.Printf("Open http://localhost:%s in the browser\n", port)
		log.Fatal(srv.ListenAndServe())
	}
}

func cleanUp() {
	Database.DisconnectDB()
}
