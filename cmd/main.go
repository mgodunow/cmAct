package main

import (
	"cmAct/internal/config"
	"log"

)

func main() {
	// TODO:need to set the waiting time for a response from the server,
	// the maximum number of requests, and if I remember something else

	a, err := config.New()
	if err != nil {
		log.Fatal(err)		
	}
	a.ListenAndServe()
}