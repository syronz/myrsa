package myrsa

import (
	"fmt"
	"math/big"
	"testing"
)

func TestGenerate(t *testing.T) {
	// SandBox()
	p := new(big.Int)
	q := new(big.Int)
	var err error

	if _, err = fmt.Sscan("2", p); err != nil {
		t.Errorf("failded in convert p to number: %v", err.Error())
	}

	if _, err = fmt.Sscan("7", q); err != nil {
		t.Errorf("failded in convert q to number: %v", err.Error())
	}

	Generate(p, q)

	// t.Log(".............")
	// fmt.Println("this is encrypt....")
}

func TestModularSolver(t *testing.T) {
	a := big.NewInt(5)
	b := big.NewInt(6)

	result := modularSolver(a, b)

	fmt.Println(a, b, result)
}
