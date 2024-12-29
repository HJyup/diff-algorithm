package diff

func longestCommonSubsequence(a, b []string) []string {
	n, m := len(a), len(b)

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if a[i] == b[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	var lcs []string
	i, j := 0, 0
	for i < n && j < m {
		if a[i] == b[j] {
			lcs = append(lcs, a[i])
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

func Diff(file1, file2 []string) []string {
	lcs := longestCommonSubsequence(file1, file2)

	var diff []string
	i, j, k := 0, 0, 0

	for i < len(file1) || j < len(file2) {
		if k < len(lcs) && i < len(file1) && j < len(file2) &&
			file1[i] == file2[j] && file1[i] == lcs[k] {
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
