# FxNGo

This is a simple Go API for obtaining FX prices from multiple providers
Currently it only support querying prices from Oanda (http://developer.oanda.com/) but there is plan to write parsers for other providers.

# How does it work?

## Code example

```go
import (
"github.com/kenshiro-o/FxNGo/model"
"github.com/kenshiro-o/FxNGo/parser"
"fmt"
)


func main(){
	oandaParser := parser.GetOandaParser()
	ccyPair := model.FxCcyPair{"EUR", "USD"}
	
	i := oandaParser.Request(ccyPair)
	prices, err := oandaParser.Parse(i)

	if err != nil{
		fmt.Println("An error occured", err)
		return
	}else{
		fmt.Println("Received and parsed FX Price", prices[0])
	}
}
```
## Data

### Currency Pair Model

```go
type FxCcyPair struct{
	Base string
	Terms string
}
```

### FX Price Model

```go
type FxPrice struct{
	CcyPair FxCcyPair
	Bid float64	
	Ask float64
	TimeStamp time.Time
}
```
