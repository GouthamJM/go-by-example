package main

import "fmt"

func assignZeroVal(val int) {
	val = 0
}

func assignZeroPointer(pointer *int) {
	*pointer = 0
}

func main() {

	i := 0
	fmt.Println("Print intial value :", i)
	assignZeroVal(i)
	fmt.Println("Print after assignZeroVal value:", i)
	assignZeroPointer(&i)
	fmt.Println("Print after assignZeroPointer value:", i)
	fmt.Println("Address of:", &i)

}
