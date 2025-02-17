package learning

import "fmt"

// we have to specify the type of the list every time
func HAs(list []string, value string) bool {

	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false

}

func HAS[T comparable](list []T, value T) bool {

	for _, v1 := range list {
		if v1 == value {
			return true
		}
	}
	return false
}
func Typeperimeters[A, B any, C ~int](a1, a2 A, b B, c C) {
	fmt.Println(a1, a2, b, c)

}

type Set[T comparable] map[T]struct{}

func Sets[T comparable](values ...T) Set[T] {
	Set := make(Set[T], len(values))
	for _, v := range values {
		Set[v] = struct{}{}
	}
	return Set
}
func (s Set[T]) Check(value T) bool {
	_, ok := s[value]
	return ok
}
