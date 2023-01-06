package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "3030"

type MyJob struct {}

func main() {
    
	// to save initial objects to the db
	initializeService()
    
	// to run auto login/auto logout/auto message functions as go routines
	mainService()

    // to generate csv log report for send messages
	logService()

	// define http server
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

