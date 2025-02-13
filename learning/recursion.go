package learning

func Recur(n int) int {
	if n == 0 {
		return 1
	}
	return n * Recur(n-1)
}
