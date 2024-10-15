package main

import (
	"fmt"
	"log"

	"github.com/KubeFinancial/bankofcanada-go/valet"
)

func main() {
	log.SetPrefix("INFO\t")
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	// Fetch the list of available series
	apiResponse, err := valet.Api("/observations/FXUSDCAD,FXEURCAD")
	if err != nil {
		log.Fatal(err)
	}

	// Print the list of groups
	log.Println(apiResponse.GroupDetail.Label)
	for name, item := range apiResponse.SeriesDetail {
		log.Printf("%s: %s\n", name, item)
	}

	for _, observation := range apiResponse.Observations {
		fmt.Println(observation)
	}
}
