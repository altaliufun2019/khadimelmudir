package main

import (
	"fmt"
	"log"
	logger "mudiralmaham/utils"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)

		logger.GeneralLogger.Printf("Defaulting to port %s\n", port)
	}
	startUpLog(false, port)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func startUpLog(inFile bool, port string) {
	if inFile {
		logger.GeneralLogger.Printf("Listening on port %s\n", port)
		logger.GeneralLogger.Printf("Open http://localhost:%s in the browser\n", port)
		logger.ErrorLogger.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	} else{
		log.Printf("Listening on port %s\n", port)
		log.Printf("Open http://localhost:%s in the browser\n", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	}
}