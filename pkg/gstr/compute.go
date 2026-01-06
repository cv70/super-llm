package gstr

// LSCLength 计算两个字符串的最长公共子序列长度
func LSCLength(s1, s2 string) int {
	m, n := len(s1), len(s2)
	// 创建一个二维数组 dp，其中 dp[i][j] 表示 s1[0:i] 和 s2[0:j] 的 LCS 长度
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 填充 dp 数组
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}
