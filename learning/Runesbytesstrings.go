package learning

import (
	"fmt"
	"unicode/utf8"
)

func Runesbytesstrings() {
	str := "Masood & Ali"

	firstletter := str[0]
	fmt.Println(firstletter)
	fmt.Println(string(firstletter))
	lastletter := str[len(str)-1]
	fmt.Println(lastletter)
	fmt.Println(string(lastletter))

	Convertrunes := []rune(str)
	fmt.Println(Convertrunes[0])
	firstletter1 := string(Convertrunes[0])
	fmt.Println((firstletter1))
	fmt.Println(Convertrunes[len(str)-1])

	lastletter1 := string(Convertrunes[len(str)-1])
	fmt.Println(lastletter1)

	for i := 0; i < len(str); i++ {
		letters := str[i]
		fmt.Println("", letters)
		letters1 := string(str[i])
		fmt.Println("", letters1)
	}
	fmt.Println("str count:", utf8.RuneCountInString(str))

	for val, val1 := range str {

		fmt.Println(val, val1)
	}

	for i, w := 0, 0; i < len(str); i += w {
		value, width := utf8.DecodeRuneInString(str[i:])

		fmt.Println(value, i)

		w = width
		runecheck(value)

	}

}

func runecheck(r rune) {

	if r == 'M' {
		fmt.Println("M is found")

	} else if r == 'A' {
		fmt.Println("A is found ")
	}
}
