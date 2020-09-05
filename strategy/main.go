package main

import (
	"fmt"
)

//strategy
type Filter interface {
	Filter([]int) []int
}

type Odd struct{}

func (*Odd) Filter(input []int) []int {
	var output []int
	for _, v := range input {
		if v%2 == 1 {
			output = append(output, v)
		}
	}
	return output
}

type Even struct{}

func (*Even) Filter(input []int) []int {
	var output []int
	for _, v := range input {
		if v%2 == 0 {
			output = append(output, v)
		}
	}
	return output
}

func main() {
	data := []int{6, 2, 14, 5, 9, 7, 11, 200}
	var filter Filter

	filter = &Odd{}
	fmt.Println(filter.Filter(data))

	filter = &Even{}
	fmt.Println(filter.Filter(data))
}
