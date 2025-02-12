// Understanding the Array
package learning

import "fmt"

func Array() {
	var a [5]int

	fmt.Println(a)

	a[4] = 100

	fmt.Println(a)

	fmt.Println(a[4])

	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	c := [3]string{"masood", "safi", "ali"}
	fmt.Println(c)

	d := [4]int{1, 3: 4}
	fmt.Println(d)

	// using 2d array
	twoda := [2][3]int{{1, 2}, {1, 2}}
	fmt.Println(twoda)

	// using 2d with for loop
	var twod [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twod[i][j] = i + j
		}
	}
	fmt.Println(twod)
}
