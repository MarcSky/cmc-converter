package internal

import (
	"context"
	"math/big"
)

type ConverterClient interface {
	Converter(amount, from, to string) (float64, error)
}

type Converter interface {
	Converter(ctx context.Context, amount *big.Float, from, to string) (*big.Float, error)
}
