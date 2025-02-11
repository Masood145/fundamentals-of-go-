package learning

import "fmt"

func Forloop() {

	//simpliest way
	i := 0
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// more pracrical way

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// using range
	for i := range 5 {
		fmt.Println(i)
	}

	// without condition

	for {
		fmt.Println("loop")
		//loop will continue if we dont use break
		break
	}
	for i := range 10 {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
