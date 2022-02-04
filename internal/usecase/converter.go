package usecase

import (
	"cmcconv/internal"
	"context"
	"math/big"
)

type conv struct {
	client internal.ConverterClient
}

func NewConv(client internal.ConverterClient) internal.Converter {
	return &conv{
		client: client,
	}
}

func (c *conv) Converter(ctx context.Context, amount *big.Float, from, to string) (*big.Float, error) {
	result, err := c.client.Convert(amount.String(), from, to)
	if err != nil {
		return nil, err
	}
	return new(big.Float).SetFloat64(result), nil
}
