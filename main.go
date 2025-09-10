package main

import (
	"fmt"
)

// LevenshteinDistance calculates the Levenshtein distance between two strings
func LevenshteinDistance(a, b string) int {
	la := len(a)
	lb := len(b)
	if la == 0 {
		return lb
	}
	if lb == 0 {
		return la
	}
	// Create a 2D slice
	dp := make([][]int, la+1)
	for i := range dp {
		dp[i] = make([]int, lb+1)
	}
	for i := 0; i <= la; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			dp[i][j] = min(
				dp[i-1][j]+1,      // deletion
				dp[i][j-1]+1,      // insertion
				dp[i-1][j-1]+cost, // substitution
			)
		}
	}
	return dp[la][lb]
}

func min(a, b, c int) int {
	m := a
	if b < m {
		m = b
	}
	if c < m {
		m = c
	}
	return m
}

func main() {
	s1 := "brah"
	s2 := "bro"
	dist := LevenshteinDistance(s1, s2)
	fmt.Printf("Levenshtein distance between '%s' and '%s' is %d\n", s1, s2, dist)

	eth1 := "0x742d32Cc6634C0535925a3b844Bc454e4438f44e"
	eth2 := "0x742d35Cc6634C0532925a3b844Bc454e4438f44f"
	dist2 := LevenshteinDistance(eth1, eth2)
	fmt.Printf("Levenshtein distance between\n'%s'\nand\n'%s'\nis %d\n\n", eth1, eth2, dist2)
}
