package app

import (
	"cmcconv/internal/cmc"
	"context"
	"math/big"
	"strconv"
	"time"
)

type App interface {
	Convert(value *big.Float, from, to string) (string, error)
}

type app struct {
	cmc cmc.CmcSvc
}

func NewApp(cmc cmc.CmcSvc) App {
	return &app{
		cmc: cmc,
	}
}

func (a *app) Convert(value *big.Float, from, to string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rateValue, err := a.cmc.Convert(ctx, value.String(), from, to)
	if err != nil {
		return "0", err
	}

	return convert(rateValue, value), nil
}

func convert(rateValue float64, value *big.Float) string {
	value.Mul(value, new(big.Float).SetFloat64(rateValue))
	res, _ := value.Float64()
	return strconv.FormatFloat(res, 'f', -1, 64)
}
