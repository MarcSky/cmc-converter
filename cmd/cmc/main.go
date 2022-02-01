package main

import (
	"cmcconv/internal/app"
	"cmcconv/internal/cmc"
	"cmcconv/internal/conv"
	"flag"
	"fmt"
)

func main() {
	var value, from, to string

	flag.StringVar(&value, "value", "0", "value")
	flag.StringVar(&from, "from", "", "from")
	flag.StringVar(&to, "to", "", "to")

	flag.Parse()

	cmcSvc := cmc.NewCMC(cmcToken)
	appl := app.NewApp(cmcSvc)
	converter := conv.NewConv(appl)

	res, err := converter.Convert(value, from, to)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
