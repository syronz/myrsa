package main

import (
	"math/big"
	"testing"
)

func TestGcd(t *testing.T) {
	a := big.NewInt(15)
	b := big.NewInt(5)

	c := Coprime(a, b)

	t.Log("test GCD has been started", c)

}
