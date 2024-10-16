package main

import (
	"fmt"
	"log"

	"github.com/KubeFinancial/bankofcanada/valet"
)

func main() {
	response, err := valet.GroupObservations("FX_RATES_DAILY")
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range response {
		fmt.Println(item)
	}
}
