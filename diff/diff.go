package diff

func longestCommonSubsequence(text1 []string, text2 []string) []string {
	if len(text1) < len(text2) {
		text1, text2 = text2, text1
	}

	dp := make([][]int, len(text2)+1)
	for i := range dp {
		dp[i] = make([]int, len(text1)+1)
	}

	for i := len(text2) - 1; i >= 0; i-- {
		for j := len(text1) - 1; j >= 0; j-- {
			if text2[i] == text1[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	var lcs []string
	i, j := 0, 0
	for i < len(text2) && j < len(text1) {
		if text2[i] == text1[j] {
			lcs = append(lcs, text2[i])
			i++
			j++
		} else if dp[i+1][j] > dp[i][j+1] {
			i++
		} else {
			j++
		}
	}

	return lcs
}

func Diff(file1 []string, file2 []string) []string {
	lcs := longestCommonSubsequence(file1, file2)

	var diff []string
	i, j, k := 0, 0, 0

	for i < len(file1) || j < len(file2) {
		if k < len(lcs) && i < len(file1) && j < len(file2) && file1[i] == file2[j] && file1[i] == lcs[k] {
			diff = append(diff, file1[i])
			i++
			j++
			k++
		} else if i < len(file1) && (k >= len(lcs) || file1[i] != lcs[k]) {
			diff = append(diff, "-"+file1[i])
			i++
		} else if j < len(file2) && (k >= len(lcs) || file2[j] != lcs[k]) {
			diff = append(diff, "+"+file2[j])
			j++
		}
	}

	return diff
}
