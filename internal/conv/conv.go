package conv

import (
	"errors"
	"math/big"

	"cmcconv/internal/app"
)

type Conv interface {
	Convert(value, from, to string) (string, error)
}

type conv struct {
	app app.App
}

func NewConv(app app.App) Conv {
	return &conv{
		app: app,
	}
}

func (c *conv) Convert(value, from, to string) (string, error) {
	v, ok := new(big.Float).SetPrec(100).SetString(value)
	if !ok {
		return "0", errors.New("wrong converting string to float")
	}

	if from == "" {
		return "0", errors.New("from cant be empty")
	}

	if to == "" {
		return "0", errors.New("to cant be empty")
	}

	return c.app.Convert(v, from, to)
}
