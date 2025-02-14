package learning

import "fmt"

type sides struct {
	length int
	width  int
}

func (r *sides) Area() int {
	r.length = r.length + 2
	r.width = r.width + 1
	return r.width * r.length
}
func (r sides) Peri() int {
	return 2*r.width + 2*r.length
}
func Result() {
	r := sides{length: 10, width: 5}
	fmt.Println(r.length)
	fmt.Println(r.width)
	fmt.Println(r.Area())
	fmt.Println(r.Peri())

	rp := &r
	fmt.Println(rp.Area())
	fmt.Println(rp.Peri())
}
