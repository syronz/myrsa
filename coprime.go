package main

import (
	"fmt"
	"math/big"
)

func gcd(a, b *big.Int) *big.Int {
	// time.Sleep(1 * time.Second)
	fmt.Println("........>>>>>>>", a, b)
	zero := big.NewInt(0)
	if a.Cmp(zero) == 0 || b.Cmp(zero) == 0 {
		return zero
	}

	if a.Cmp(b) == 0 {
		return a
	}

	if a.Cmp(b) == 1 {
		bNeg := new(big.Int)
		*bNeg = *b
		bNeg.Neg(bNeg)
		return gcd(a.Add(a, bNeg), b)
	}

	// return zero
	aNeg := new(big.Int)
	*aNeg = *a
	aNeg.Neg(aNeg)
	return gcd(a, b.Add(b, aNeg))
}

// Coprime is checking two numbers are coprime
func Coprime(a, b *big.Int) bool {
	one := big.NewInt(1)
	result := gcd(a, b)
	if result.Cmp(one) == 0 {
		return true
	}

	return false
}

/*
function __gcd($a, $b)
    {
        // Everything divides 0
        if ($a == 0 || $b == 0)
            return 0;

        // base case
        if ($a == $b)
            return $a;

        // a is greater
        if ($a > $b)
            return __gcd($a - $b, $b);

        return __gcd($a, $b - $a);
    }

    // function to check and print if
    // two numbers are co-prime or not
function coprime($a, $b)
{
    if (__gcd($a, $b) == 1)
        echo "Co-Prime","\n";
    else
        echo "Not Co-Prime","\n";
}
*/
