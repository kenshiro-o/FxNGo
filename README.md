# FxNGo

This is a simple Go API for obtaining FX prices from multiple providers
Currently it only support querying prices from Oanda (http://developer.oanda.com/) but there is plan to write parsers for other providers.

## How does it work?

### Code example

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
### Data

#### Currency Pair Model

```go
type FxCcyPair struct{
	Base string
	Terms string
}
```

#### FX Price Model

```go
type FxPrice struct{
	CcyPair FxCcyPair
	Bid float64	
	Ask float64
	TimeStamp time.Time
}
```


## Licence

(The MIT License)

Copyright (c) 2013 Kenshiro &lt;kenshiro@kenshiro.me&gt;

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
'Software'), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
