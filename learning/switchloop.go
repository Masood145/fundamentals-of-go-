package learning

//using time package and its methods in switch for practice
import (
	"fmt"
	"time"
)

func Switch() {
	// simple switch
	i := 8
	switch i {
	case 6:
		fmt.Println(" i is 6 ")
	case 8:
		fmt.Println("i is 8  ")
	default:
		fmt.Println("its not a number ")
	}
	// using switch when conditon is used in case statement
	num := 9
	switch {
	case num%2 == 0:
		fmt.Println(num, "its even ")
	case num&2 != 0:
		fmt.Println(num, " ist odd ")
	default:
		fmt.Println("its not a number ")
	}
	// switch using time package
	t := time.Now()
	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("day is weekend ")
	default:
		fmt.Println("dai is workday ")
	}
	//another example using time
	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon ")
	case t.Hour() > 12:
		fmt.Println("its after noon ")
	}
	// check the type of interface value

	whatistype := func(typ interface{}) {
		switch v := typ.(type) {
		case bool:
			fmt.Println(v, "its bool type ")
		case int:
			fmt.Println("its int type ")
		case string:
			fmt.Println("its string type")
		default:
			fmt.Println("its is Null")
		}

	}
	whatistype(9)
	whatistype(true)
	whatistype("hello")

}
