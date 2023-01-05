package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

const webPort = "3030"

type MyJob struct {}
var wait sync.WaitGroup

func main() {

	log.Println("Start test tool")
    
	// to save initial objects to the db
	initializeService()
    
	// to run auto login/auto logout/auto message functions as go routines
	mainService()

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

