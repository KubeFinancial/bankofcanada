package main

import (
	"log"
	"os"

	"github.com/KubeFinancial/bankofcanada/valet"
)

func main() {
	logger := log.New(os.Stdout, "", 0)

	displaySeriesAndGroupsCount(logger)
	displaySeriesInfo(logger)
	displayGroupInfo(logger)
	displaySeriesObservations(logger)
	displayGroupObservations(logger)
}

func displaySeriesAndGroupsCount(logger *log.Logger) {
	seriesList, err := valet.ListSeries()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("Number of series: %d", len(seriesList))

	groupsList, err := valet.ListGroups()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("Number of groups: %d", len(groupsList))
}

func displaySeriesInfo(logger *log.Logger) {
	series, err := valet.Series("FXUSDCAD")
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf(
		"%s\t%s\t%s",
		series.Dimension.Name,
		series.Label,
		series.Description,
	)
}

func displayGroupInfo(logger *log.Logger) {
	group, err := valet.Group("FX_RATES_DAILY")
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf(
		"%s\t%s",
		group.Detail.Name,
		group.Detail.Description,
	)

	for _, groupSeries := range group.GroupSeries {
		logger.Printf(
			"%s\t%s\t%s",
			groupSeries.Dimension.Name,
			groupSeries.Label,
			groupSeries.Description,
		)
	}
}

func displaySeriesObservations(logger *log.Logger) {
	seriesObservations, err := valet.SeriesObservations("FXUSDCAD,FXEURCAD")
	if err != nil {
		logger.Fatal(err)
	}

	for _, seriesObservation := range seriesObservations {
		logger.Printf(
			"%s\t%s\t%s",
			seriesObservation.Date,
			seriesObservation.Name,
			seriesObservation.Value,
		)
	}
}

func displayGroupObservations(logger *log.Logger) {
	groupObservations, err := valet.GroupObservations("FX_RATES_DAILY")
	if err != nil {
		logger.Fatal(err)
	}

	for _, seriesObservation := range groupObservations {
		logger.Printf(
			"%s\t%s\t%s",
			seriesObservation.Date,
			seriesObservation.Name,
			seriesObservation.Value,
		)
	}
}
