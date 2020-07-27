# client_covid
Golang HTTP client for https://github.com/junkd0g/covid

## How to use

### Instalation

go get "github.com/junkd0g/client_covid"

### Usage

```
package main

import (
	"fmt"

	jkdcovid "github.com/junkd0g/client_covid"
)

func main() {

	csseData, cssDataError := jkdcovid.GetCSSEData("USA")
	if cssDataError != nil {
		fmt.Println(cssDataError.Error())
	}
	fmt.Println(csseData)

}
```

Check examples/main.go for more examples