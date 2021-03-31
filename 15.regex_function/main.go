package main

import (
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("d([a-z])+gn", "design")
	fmt.Println(match)

	r, _ := regexp.Compile("d([a-z])+gn")

	fmt.Println(r.MatchString("degn"))
	fmt.Println(r.FindString("design"))
}
