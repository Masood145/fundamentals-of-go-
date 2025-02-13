package learning

func Clouser() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
