package main

import (
	"fmt"
	"log"

	"github.com/KubeFinancial/bankofcanada/valet"
)

func main() {
	series, err := valet.Series("FXUSDCAD")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(series)

	group, err := valet.Group("FX_RATES_DAILY")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(group)

	groupObservations, err := valet.GroupObservations("FX_RATES_DAILY")
	if err != nil {
		log.Fatal(err)
	}
	for _, seriesObservation := range groupObservations {
		fmt.Println(seriesObservation)
	}
}
