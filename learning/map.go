package learning

import (
	"fmt"
	"maps"
)

func Map() {
	//declare map
	m := make(map[string]int)
	//initilize map
	m["key1"] = 5
	m["key2"] = 9
	fmt.Println(m)
	//assign values of map to aniother variable
	v1 := m["key1"]
	v2 := m["key3"]

	fmt.Println(v1)
	fmt.Println(v2)
	//deleting the key from the map
	delete(m, "key1")
	fmt.Println(m)
	//Clearing the key
	clear(m)
	fmt.Println(m)
	//checking the key present in the map
	_, check := m["key1"]
	fmt.Println(check)
	//  Declaring and initilizing the mao in one line
	newmap := map[string]int{"key1": 1, "key2": 4}
	fmt.Println(newmap)
	//  using the map method
	//newmap1 := map[string]int{"key1": 1, "key2": 4}
	// if maps.Equal(newmap,newmap1)
	if maps.Equal(newmap, m) {
		fmt.Println("is equal ")
	} else {
		fmt.Println("not equal")
	}

}
