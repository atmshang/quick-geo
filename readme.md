
# quick-geo

## Overview

`quick-geo` is a Go library that helps you find the corresponding geo information given a specific longitude and latitude. It leverages the `orb` package to handle geometric types and operations in Go.

## Installation

To install `quick-geo`, you can use `go get`:

```bash
go get github.com/atmshang/quick-geo
```

## Usage

Here is a basic example of how to use `quick-geo`:

```go
package main

import (
	"fmt"
	"github.com/atmshang/quick-geo"
)

func main() {
	name := quick_geo.GetCountryOrRegionName(114.03, 22.61)
	fmt.Println(name)

	name = quick_geo.GetCountryOrRegionName(122.96, 12.98)
	fmt.Println(name)
}
```

In this example, `GetCountryOrRegionName` is used to get the name of the country or region for a given longitude and latitude.

## License

This project is licensed under the MIT License.