package learning

import "fmt"

func Rangeovertypes() {

	arr := []int{2, 4, 5, 7, 8}
	sum := 0
	for _, arr1 := range arr {
		sum += arr1
	}

	fmt.Println(sum)

	for i, arr2 := range arr {
		if arr2 == 7 {
			fmt.Println("index ", i)
		}

	}

	Fruits := map[int]string{
		1: "apple",
		2: "mango",
	}
	for f1, f2 := range Fruits {

		fmt.Println(f1, f2)
	}
	for k := range arr {
		fmt.Println(k)

	}
}
