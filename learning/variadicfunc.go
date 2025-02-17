package learning

import "fmt"

func Varfunc(numbers ...int) {
	fmt.Print(numbers)
	total := 0

	for _, sum := range numbers {
		total += sum
	}
	fmt.Println(total)
}
