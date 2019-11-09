package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	numOfDigits, err := strconv.Atoi(os.Args[1])
	if err != nil || numOfDigits <= 0 {
		fmt.Println("First argument should be # of digits as an positive integer")
		os.Exit(1)
	}
	var lowBound, highBound uint64

	lowBound = uint64(math.Pow10(numOfDigits - 1))
	highBound = uint64(math.Pow10(numOfDigits)) - 1
	fmt.Printf("Lowbound %d; Highbound %d \n", lowBound, highBound)
	for n := highBound * highBound; n >= lowBound*lowBound; n-- {
		if isPalindrome(n) {
			//Uncomment below line if all palindromes are to be printed
			//fmt.Printf("%d is a palindrome\n", n)
			factorable, fact1, fact2 := hasProperFactors(lowBound, highBound, n)
			if factorable == true {
				fmt.Printf("Largest palindrome as a product of two %d-digit numbers is %d = %d * %d", numOfDigits, n, fact1, fact2)
				return
			}
		}
	}
	fmt.Println("Palindrome does not exist")
}

func isPalindrome(origNum uint64) bool {
	var r uint64
	n := origNum
	for n != 0 {
		r = r * 10
		r = r + n%10
		n = n / 10
	}
	return origNum == r
}

func hasProperFactors(lowBound, highBound, n uint64) (bool, uint64, uint64) {
	for i := highBound; i >= lowBound; i-- {
		if n%i == 0 {
			o := n / i
			if o >= lowBound && o <= highBound {
				return true, i, o
			}
		}
	}
	return false, 0, 0
}
