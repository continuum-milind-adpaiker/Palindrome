# Palindrome Finder

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.Find the largest palindrome made from the product of two 3-digit numbers.

Programm is generic in nature so that it works for any "n" digit numbers. If you want to solve it specifically for 3-digit numbers it can be further optimized. 
Compile and run this programm as follows by providing umber of digits as os argument
  Milind.Adpaiker@RMM-PN-LT-092 /cygdrive/d/gowork/src/github.com/milind/palindrome
  $ go build -o palindromeFinder

  Milind.Adpaiker@RMM-PN-LT-092 /cygdrive/d/gowork/src/github.com/milind/palindrome
  $ ./palindromeFinder 3
  Lowbound 100; Highbound 999
  Largest palindrome as a product of two 3-digit numbers is 906609 = 993 * 913

Go playground friedly version here: https://play.golang.org/p/WMpOv0oHACm Change number of digits at Line #9
