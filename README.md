# Unofficial Bank of Canada Valet API Go Module

This is an **unofficial Go module** for interacting with the [Bank of Canada Valet API](https://www.bankofcanada.ca/valet/docs). The module allows developers to programmatically access a wide range of economic data including exchange rates, interest rates, and financial time series published by the Bank of Canada.

## Features

- Query and fetch the latest economic data from the Bank of Canada.
- Easily integrate financial data into Go applications.

## Prerequisites

- **Go 1.22 or higher**

## Installation

Install the module by running:

```bash
go get -u github.com/KubeFinancial/bankofcanada
```

## Usage

```go
package main

import (
	"log"
	"os"

	"github.com/KubeFinancial/bankofcanada/valet"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	
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


```

## Example Output
The above code will return the most recent FX_RATES_DAILY data from the Bank of Canada Valet API.
```shell
2024/10/22 06:42:27.920229 Request: GET https://www.bankofcanada.ca/valet/observations/group/FX_RATES_DAILY/json?recent=1
2024/10/22 06:42:27.935205 Response: 200 OK, Filename: FX_RATES_DAILY.json, Generated: 2024-10-22 10:38:26 UTC
2019-12-31      FXVNDCAD        0.000056
2019-12-31      FXMYRCAD        0.3175
2019-12-31      FXTHBCAD        0.04362
2024-10-21      FXTRYCAD        0.04040
2024-10-21      FXIDRCAD        0.000089
2024-10-21      FXZARCAD        0.07854
2024-10-21      FXSEKCAD        0.1312
2024-10-21      FXUSDCAD        1.3835
2024-10-21      FXEURCAD        1.4983
2024-10-21      FXKRWCAD        0.001003
2024-10-21      FXHKDCAD        0.1780
2024-10-21      FXCNYCAD        0.1944
2024-10-21      FXSGDCAD        1.0523
2024-10-21      FXINRCAD        0.01646
2024-10-21      FXGBPCAD        1.7983
2024-10-21      FXAUDCAD        0.9228
2024-10-21      FXMXNCAD        0.06918
2024-10-21      FXNOKCAD        0.1265
2024-10-21      FXCHFCAD        1.5991
2024-10-21      FXNZDCAD        0.8358
2024-10-21      FXRUBCAD        0.01431
2024-10-21      FXJPYCAD        0.009200
2024-10-21      FXTWDCAD        0.04312
2024-10-21      FXBRLCAD        0.2424
2024-10-21      FXPENCAD        0.3678
2024-10-21      FXSARCAD        0.3683

```

## Additional Endpoints
Refer to the official [Valet API Documentation](https://www.bankofcanada.ca/valet/docs)  for more available endpoints and options.

## License and Attribution

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

You are free to use, modify, and distribute the code as needed. However, please note the following attribution:

Source: [Bank of Canada](https://www.bankofcanada.ca/terms/)

Content has been modified from its original form. This project is not endorsed by or affiliated with the Bank of Canada.
