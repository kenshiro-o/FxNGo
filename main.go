package main

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
