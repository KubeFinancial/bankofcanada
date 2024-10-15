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
	apiResponse, err := valet.Api("/observations/group/FX_RATES_DAILY/json?recent=2")
	if err != nil {
		log.Fatal(err)
	}

	// Print the unmarshalled data
	for _, observation := range apiResponse.Observations {
		for pair, rate := range observation.Series {
			fmt.Println(observation.Date, pair, rate.Value)
		}
	}
}
