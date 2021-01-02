package services

import (
	"fmt"
)

type StocksService struct{}

func NewStocksService() *StocksService {
	return &StocksService{}
}

func (a *StocksService) Test() {
	fmt.Println("hello")
}
