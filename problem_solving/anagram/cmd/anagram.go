//Given two strings s and p, return an array of all the start indices of p's anagrams in s.
//You may return the answer in any order.

package main

import "fmt"

func main() {
	anagrams := findAnagrams("cbaebabacdjkl", "jkl")
	fmt.Println(anagrams)
}

func findAnagrams(s string, p string) []int {
	ana := map[byte]int{}
	length := len(p)
	ans := []int{}

	for _, r := range []byte(p) {
		ana[r]++
	}

	for idx, r := range []byte(s) {
		ana[r]--
		if ana[r] == 0 {
			delete(ana, r)
		}
		if idx-length >= 0 {
			ana[s[idx-length]]++
			if ana[s[idx-length]] == 0 {
				delete(ana, s[idx-length])
			}
		}
		if len(ana) == 0 {
			ans = append(ans, idx-length+1)
		}
	}
	return ans
}
