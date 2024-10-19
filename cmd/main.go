package main

import (
	"fmt"
	"log"

	"github.com/KubeFinancial/bankofcanada/valet"
)

func main() {
	series, err := valet.AwesomeSeries("FXUSDCAD")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(series)

	groupObservations, err := valet.GroupObservations("FX_RATES_DAILY")
	if err != nil {
		log.Fatal(err)
	}
	for _, seriesObservation := range groupObservations {
		fmt.Println(seriesObservation)
	}
}
