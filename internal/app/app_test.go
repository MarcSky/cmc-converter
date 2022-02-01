package app

import (
	"math/big"
	"testing"
)

func TestApp_Convert(t *testing.T) {
	type test struct {
		a   float64
		b   *big.Float
		res string
	}

	var tests = []test{
		{
			a:   0.0000001,
			b:   new(big.Float).SetPrec(100).SetFloat64(1000000),
			res: "0.09999999999999999",
		},
		{
			a:   0.00000000001,
			b:   new(big.Float).SetPrec(100).SetFloat64(1235),
			res: "0.00000001235",
		},
		{
			a:   1294,
			b:   new(big.Float).SetPrec(100).SetFloat64(1000000),
			res: "1294000000",
		},
	}

	for i := range tests {
		res := convert(tests[i].a, tests[i].b)
		if res != tests[i].res {
			t.Errorf("must be %s but we got %s", tests[i].res, res)
		}
	}
}
