package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2 * (r.width + r.height)
}

func main() {

	r := rect{
		width: 10, height: 20,
	}

	fmt.Println(r.area(), "area")
	fmt.Println(r.perim(), "perimeter")

	rp := &r
	fmt.Println(rp.area(), "area")
	fmt.Println(rp.perim(), "perimeter")
}
