package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

func Generate(p, q *big.Int) (err error) {

	fmt.Println("p & q", p, q)

	pDec := new(big.Int)
	qDec := new(big.Int)

	pDec.Add(p, big.NewInt(-1))
	qDec.Add(q, big.NewInt(-1))

	N := new(big.Int)
	N.Mul(p, q)

	PHI := new(big.Int)
	PHI.Mul(pDec, qDec)

	fmt.Println("p & q after add", p, q, N, pDec, qDec, PHI)

	return
}

func SandBox() {
	fmt.Println(":::::: - ")

	// a := big.NewInt(0)
	b := big.NewInt(1)
	// // b := big.NewInt(112345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567)
	// c, d, err := big.ParseFloat("99999", 10, 10, big.ToZero)

	// fmt.Println(a, b, c, d, err)
	// a.Add(a, b)
	// fmt.Println(a, b)

	f, err := os.Open("primes/list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	i := new(big.Int)
	for scanner.Scan() {
		numStr := scanner.Text()
		fmt.Println(numStr)
		fmt.Sscan(numStr, i)
	}

	// _, err = fmt.Sscan("112345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567", i)
	i.Add(i, b)
	fmt.Println("@@@@@@@@@@", i, err)

	now := time.Now()
	fmt.Println(now)
	two := big.NewInt(10)

	b = i.Mod(i, two)
	fmt.Println(time.Since(now))
	fmt.Println(">>>>>>>>>> ", b, two)
}
