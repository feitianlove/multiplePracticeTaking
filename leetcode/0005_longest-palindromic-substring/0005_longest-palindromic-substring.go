package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome2("babad"))
}
func longestPalindrome(s string) string {
	max := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			// 所有字串 s[i:j]
			if isPalindrome(s[i:j]) == true {
				if len(max) < len(s[i:j]) {
					max = s[i:j]
				}
			}
		}
	}
	return max
}
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

//动态规划
func longestPalindrome2(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, len(s))
	start := 0
	max := 1
	for r := 0; r < len(s); r++ {
		dp[r] = make([]bool, len(s))
		dp[r][r] = true
		fmt.Println("start", dp, r)
		for l := 0; l < r; l++ {
			if s[l] == s[r] && (r-l <= 2 || dp[l+1][r-1] == true) {
				dp[l][r] = true
			} else {
				dp[l][r] = false
			}
			if dp[l][r] == true {
				if r-l+1 > max {
					max = r - l + 1
					start = l
				}
			}

		}
		fmt.Println("end", dp)
	}
	return s[start : start+max]
}
