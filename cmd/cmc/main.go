package main

import (
	"cmcconv/internal/cmc"
	"cmcconv/internal/usecase"
	"context"
	"errors"
	"flag"
	"fmt"
	"math/big"
	"strconv"
)

const (
	rate float64 = 15
)

func main() {
	var value, from, to string

	flag.StringVar(&value, "value", "0", "value")
	flag.StringVar(&from, "from", "", "from")
	flag.StringVar(&to, "to", "", "to")
	flag.Parse()

	v, ok := new(big.Float).SetPrec(100).SetString(value)
	if !ok {
		panic(errors.New("wrong converting string to float"))
	}

	if from == "" {
		panic(errors.New("from cant be empty"))
	}

	if to == "" {
		panic(errors.New("to cant be empty"))
	}

	cmcSvc := cmc.NewCMC(cmcToken)
	converter := usecase.NewConv(cmcSvc)

	converterWithInterest := usecase.NewConvWithInterest(converter, rate)
	res, err := converterWithInterest.Converter(context.Background(), v, from, to)
	if err != nil {
		panic(err)
	}

	r, _ := res.Float64()
	fmt.Println(strconv.FormatFloat(r, 'f', -1, 64))
}
