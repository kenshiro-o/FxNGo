package parser

import(
 	"testing"
	"github.com/kenshiro-o/FxNGo/model"
 )


func TestParseEURUSD(t *testing.T){
	oandaParser := GetOandaParser()
	ccyPair := model.FxCcyPair{"EUR", "USD"}

	i := oandaParser.Request(ccyPair)
	prices, err := oandaParser.Parse(i)

	if err != nil{
		t.Errorf("An error occurred when trying to get price [ccypair=%v, error=%v]", ccyPair, err)
	}

	if len(prices) != 1{
		t.Errorf("Length of prices array is invdalid [expected=1, actual=%d]", len(prices))
	}
}