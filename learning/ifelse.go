package learning

import "fmt"

func Ifelse() {
	if 6%2 == 0 {
		fmt.Println("6 is even ")
	} else {
		fmt.Println("6 bis odd ")
	}

	if 6 > 1 {
		fmt.Println("6 is greater than one ")
	}
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("8 or 7 is even ")
	}
	if num := 9; num > 6 {
		fmt.Println(num, "is less than 6 ")
	} else if num%2 == 0 {
		fmt.Println(num, "it is even number ")
	} else {
		fmt.Println(num, "is not a digit ")
	}
}
