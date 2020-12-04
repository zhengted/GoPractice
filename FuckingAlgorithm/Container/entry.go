package main

import (
	"GoPractice/FuckingAlgorithm/Container/Vector"
	"fmt"
)

func CreateVector() Vector.Vector {
	s := []interface{}{}
	v := Vector.Vector{
		S: s,
	}
	return v
}

func main() {

	v := CreateVector()
	v.Push_back(5)
	v.Push_back(3)
	v.Erase(0)
	fmt.Println(v)

}
