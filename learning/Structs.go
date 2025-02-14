package learning

import "fmt"

type man struct {
	name string
	age  int
}

func newman(name string) *man {
	m := man{name: name}
	m.age = 56
	return &m
}
func Struc() {
	fmt.Println(man{"masood", 78})
	fmt.Println(man{name: "masood "})
	fmt.Println(man{name: "masood ", age: 76})
	fmt.Println(&man{"masood", 98})
	fmt.Println(newman("masood"))

	new := man{name: "ali", age: 98}

	new1 := &new
	fmt.Println(new1.age)
	fmt.Println(new1.name)
	new.age = 45
	fmt.Println(new.age)

	animal := struct {
		name string
		kind string
	}{
		"dog",
		"german",
	}
	fmt.Println(animal)

}
