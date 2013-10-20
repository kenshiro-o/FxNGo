package model

import "time"

//Basic structure for a currency pair
type FxCcyPair struct{
	Base string
	Terms string
}

//Basic structure for an Fx price
type FxPrice struct{
	CcyPair FxCcyPair
	Bid float64	
	Ask float64
	TimeStamp time.Time
}