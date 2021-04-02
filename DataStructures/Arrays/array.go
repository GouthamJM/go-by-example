package array

// Creating a Structhre for Array of type Int
type Array struct {
	data   []int
	length int
}

// function to Initialize an Array
func NewArray(capacity uint) *Array {
	if capacity == 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity),
		length: 0,
	}
}

// method on Array to find it's length
func (a *Array) Len() int {
	return a.length
}

//

func (a *Array) isIndexOutofRange(index uint) bool {
	intIndex := int(index)
	return intIndex >= a.length
}
