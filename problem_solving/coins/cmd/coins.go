/*
You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where the ith row
has exactly i coins. The last row of the staircase may be incomplete.

Given the integer n, return the number of complete rows of the staircase you will build.
*/
package main

import "fmt"

func main() {
	fmt.Printf("%v", arrangeCoins(5))
}

func arrangeCoins(n int) int {
	cnt := 1
	var sum int
	for {
		sum += cnt
		if n < sum {
			return cnt - 1
		} else if n == sum {
			return cnt
		}
		cnt += 1
	}
}
