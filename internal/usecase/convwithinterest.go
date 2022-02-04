package usecase

import (
	"cmcconv/internal"
	"context"
	"math/big"
)

type ConvWithInterest struct {
	u    internal.Converter
	rate float64
}

func NewConvWithInterest(u internal.Converter, rate float64) *ConvWithInterest {
	return &ConvWithInterest{
		u: u, rate: rate,
	}
}

func (c *ConvWithInterest) Converter(ctx context.Context, amount *big.Float, from, to string) (*big.Float, error) {
	rateValue, err := c.u.Converter(ctx, amount, from, to)
	if err != nil {
		return nil, err
	}
	r := rateValue.Mul(rateValue, new(big.Float).SetFloat64(c.rate))
	return r, nil
}
