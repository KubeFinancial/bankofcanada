# Unofficial Bank of Canada Valet API Go Module

[![Go Reference](https://pkg.go.dev/badge/github.com/KubeFinancial/bankofcanada.svg)](https://pkg.go.dev/github.com/KubeFinancial/bankofcanada)
[![Go Version](https://img.shields.io/github/go-mod/go-version/KubeFinancial/bankofcanada)](https://golang.org/doc/devel/release.html)
[![License](https://img.shields.io/github/license/KubeFinancial/bankofcanada)](https://github.com/KubeFinancial/bankofcanada/blob/main/LICENSE)

This is an **unofficial Go module** for interacting with the [Bank of Canada Valet API](https://www.bankofcanada.ca/valet/docs). The module allows developers to programmatically access a wide range of economic data including exchange rates, interest rates, and financial time series published by the Bank of Canada.

## Features

- Query and fetch the latest economic data from the Bank of Canada.
- Easily integrate financial data into Go applications.
- Support for various API endpoints including series, groups, and observations.

## Prerequisites

- **Go 1.22 or higher**

## Installation

Install the module by running:

```bash
go get -u github.com/KubeFinancial/bankofcanada
```

## Usage

Here's a basic example of how to use the module with the API function:

```go
package main

import (
    "log"
    "os"

    "github.com/KubeFinancial/bankofcanada/valet"
)

func main() {
    logger := log.New(os.Stdout, "", 0)

    endpointURL := "https://www.bankofcanada.ca/valet/observations/group/FX_RATES_DAILY/json?recent=1"

    response, err := valet.API(endpointURL)
    if err != nil {
        logger.Fatal(err)
    }

    for _, observation := range response.Observations {
        for seriesName, seriesObs := range observation.Series {
            logger.Printf(
                "%s\t%s\t%s",
                observation.Date,
                seriesName,
                seriesObs.Value,
            )
        }
    }
}
```

### More Examples

#### Listing All Series

```go
seriesList, err := valet.ListSeries()
if err != nil {
    log.Fatal(err)
}
for name, details := range seriesList {
    log.Printf("Series: %s, Label: %s", name, details.Label)
}
```

#### Fetching Series Information

```go
series, err := valet.Series("FXUSDCAD")
if err != nil {
    log.Fatal(err)
}
log.Printf("Series: %s, Description: %s", series.Label, series.Description)
```

#### Fetching Group Observations

```go
groupObservations, err := valet.GroupObservations("FX_RATES_DAILY")
if err != nil {
    log.Fatal(err)
}
for _, seriesObservation := range groupObservations {
    log.Printf(
        "%s\t%s\t%s",
        seriesObservation.Date,
        seriesObservation.Name,
        seriesObservation.Value,
    )
}
```

#### Fetching Series Observations

```go
seriesObservations, err := valet.SeriesObservations("FXUSDCAD,FXEURCAD")
if err != nil {
    log.Fatal(err)
}
for _, seriesObservation := range seriesObservations {
    log.Printf(
        "%s\t%s\t%s",
        seriesObservation.Date,
        seriesObservation.Name,
        seriesObservation.Value,
    )
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License and Attribution

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

You are free to use, modify, and distribute the code as needed. However, please note the following attribution:

Source: [Bank of Canada](https://www.bankofcanada.ca/terms/)

Content has been modified from its original form. This project is not endorsed by or affiliated with the Bank of Canada.
