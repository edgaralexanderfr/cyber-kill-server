package util

func CompareEpsilon(a, b, epsilon float32) bool {
	return a >= b-epsilon && a <= b+epsilon
}
