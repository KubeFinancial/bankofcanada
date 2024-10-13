package main

import (
	"log"

	"github.com/KubeFinancial/bankofcanada-go/valet"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	// Fetch the list of available series
	response, err := valet.Api("/lists/series/json")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(response)
}
