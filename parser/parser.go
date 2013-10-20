package parser

import "github.com/kenshiro-o/FxNGo/model"

type FxParser interface{
	//Request data for this currency pair
	Request(model.FxCcyPair) *interface{}
	//Parse the received data and return an FxPrice
	Parse(*interface{}) ([] model.FxPrice, error)	
}

//Returns an oanda parser
func GetOandaParser() FxParser{
	p := new(oandaFxParser)
	return p
}