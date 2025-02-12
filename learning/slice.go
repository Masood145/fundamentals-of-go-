package learning

import (
	"fmt"
	"slices"
)

func Slice() {
	// nutt silce
	var s []int
	fmt.Println(s, len(s), cap(s))
	// slice with index 3 prinitng slice , length , capacity
	var s1 = make([]int, 3)
	fmt.Println(s1, len(s1), cap(s1))
	// adding values to indexs
	s1[0] = 1
	s1[2] = 2
	s1[1] = 3
	// printing after after the values at indexes
	fmt.Println(s1)
	fmt.Println(s1[2])
	// adding more values to null  slice
	s = append(s, 2)
	fmt.Println(s)
	// adding more values to slice s1
	s1 = append(s1, 3)
	s1 = append(s1, 5, 7)
	fmt.Println(s1)
	//making copy of slice
	Copy := make([]int, len(s1))
	copy(Copy, s1)
	fmt.Println(Copy)

	// now moving to slicing operator

	l := s1[2:5]
	fmt.Println(l)

	l1 := s1[:5]
	fmt.Println(l1)

	l2 := s1[2:]
	fmt.Println(l2)

	newslice := []string{"masood", "safi", "ali"}
	fmt.Println(newslice)

	//comparing two slice fro equal thay should have same things in their indexes

	newslice1 := []string{"masood", "safi", "ali"}
	if slices.Equal(newslice, newslice1) {
		fmt.Println("is equal ")
	} else {
		fmt.Println("not equal ")
	}

	twod := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlength := i + 1
		twod[i] = make([]int, innerlength)
		for j := 0; j < innerlength; j++ {
			twod[i][j] = i + j

		}
	}
	fmt.Println(twod)

}
