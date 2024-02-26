package main

import "fmt"

// There may be multiple longest common subsequences, but this solution only returns one

func lcsAlgo(S1, S2 string) string {
	m, n := len(S1), len(S2)
	LCS_table := make([][]int, m+1)
	for i := range LCS_table {
		LCS_table[i] = make([]int, n+1)
	}

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				LCS_table[i][j] = 0
			} else if S1[i-1] == S2[j-1] {
				LCS_table[i][j] = LCS_table[i-1][j-1] + 1
			} else {
				LCS_table[i][j] = max(LCS_table[i-1][j], LCS_table[i][j-1])
			}
		}
	}

	index := LCS_table[m][n]
	lcs := make([]byte, index)

	i, j := m, n
	for i > 0 && j > 0 {
		if S1[i-1] == S2[j-1] {
			lcs[index-1] = S1[i-1]
			i--
			j--
			index--
		} else if LCS_table[i-1][j] > LCS_table[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return string(lcs)
}

func main() {
	S1 := "ACADB"
	S2 := "CBDAD"
	LCS := lcsAlgo(S1, S2)
	fmt.Println("S1:", S1)
	fmt.Println("S2:", S2)
	fmt.Println("LCS:", LCS)
}
