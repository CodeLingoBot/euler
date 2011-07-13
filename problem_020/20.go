package main

import (
	"big"
	"fmt"
)

func SumDigitsOf(number *big.Int) int64 {
	digits := number.String()
	sum := int64(0)
	for _, digit := range digits {
		sum += int64(digit - '0')
	}

	return sum
}

func Factorial(number int64) *big.Int {
	factorial := big.NewInt(1)

	for i := number; i >= 1; i-- {
		scale := big.NewInt(i)
		factorial.Mul(factorial, scale)
	}

	return factorial
}

func main() {
	fact := Factorial(100)
	fmt.Println(SumDigitsOf(fact))
}
