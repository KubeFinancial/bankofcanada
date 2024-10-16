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

```

## Example Output
The above code will return the most recent FX_RATES_DAILY data from the Bank of Canada Valet API.
```shell
{"severity":"DEBUG","timestamp":"2024-10-16T03:57:28.032-0400","message":"Request: GET https://www.bankofcanada.ca/valet/observations/group/FX_RATES_DAILY/json?recent=1"}
{"severity":"DEBUG","timestamp":"2024-10-16T03:57:28.299-0400","message":"Response: 200 OK, Filename: FX_RATES_DAILY.json, Generated: 2024-10-16 06:01:00 UTC"}
{2019-12-31  FXVNDCAD 0.000056}
{2019-12-31  FXMYRCAD 0.3175}
{2019-12-31  FXTHBCAD 0.04362}
{2024-10-15  FXGBPCAD 1.8055}
{2024-10-15  FXNZDCAD 0.8403}
{2024-10-15  FXAUDCAD 0.9262}
{2024-10-15  FXIDRCAD 0.000089}
{2024-10-15  FXTWDCAD 0.04287}
{2024-10-15  FXTRYCAD 0.04030}
{2024-10-15  FXPENCAD 0.3666}
{2024-10-15  FXSEKCAD 0.1329}
{2024-10-15  FXUSDCAD 1.3805}
{2024-10-15  FXSARCAD 0.3677}
{2024-10-15  FXEURCAD 1.5044}
{2024-10-15  FXNOKCAD 0.1276}
{2024-10-15  FXSGDCAD 1.0545}
{2024-10-15  FXKRWCAD 0.001013}
{2024-10-15  FXCHFCAD 1.6010}
{2024-10-15  FXZARCAD 0.07830}
{2024-10-15  FXHKDCAD 0.1777}
{2024-10-15  FXINRCAD 0.01642}
{2024-10-15  FXMXNCAD 0.07033}
{2024-10-15  FXBRLCAD 0.2446}
{2024-10-15  FXCNYCAD 0.1939}
{2024-10-15  FXJPYCAD 0.009250}
{2024-10-15  FXRUBCAD 0.01428}
```

## Additional Endpoints
Refer to the official [Valet API Documentation](https://www.bankofcanada.ca/valet/docs)  for more available endpoints and options.

## License and Attribution

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

You are free to use, modify, and distribute the code as needed. However, please note the following attribution:

Source: [Bank of Canada](https://www.bankofcanada.ca/terms/)

Content has been modified from its original form. This project is not endorsed by or affiliated with the Bank of Canada.
