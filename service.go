package main 

import (
	"context"
	"fmt"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)

}

type priceFetcher struct {}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error)
{

}


var priceMocks = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
	"GG": 100_000.0,
}


func MockPriceFetcher (ctx context.Context, ticker string )(float64, error){

}