package myrsa

import (
	"math/big"
	"testing"
)

func TestCoprime(t *testing.T) {
	samples := []struct {
		a      *big.Int
		b      *big.Int
		result bool
	}{
		{
			a:      big.NewInt(15),
			b:      big.NewInt(5),
			result: false,
		},
		{
			a:      big.NewInt(15),
			b:      big.NewInt(7),
			result: true,
		},
	}

	for _, v := range samples {
		result := Coprime(v.a, v.b)
		if result != v.result {
			t.Errorf("Coprime(%v,%v) should be %v but it is %v", v.a, v.b, v.result, result)
		}
	}

}
