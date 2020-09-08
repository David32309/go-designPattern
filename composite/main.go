package main

import (
	"fmt"
)

type Condition interface {
	Match(i int) bool
}

type Even struct{}

func (*Even) Match(i int) bool {
	return i%2 == 0
}

type Greater struct {
	param int
}

func (g *Greater) Match(i int) bool {
	return i >= g.param
}

type And struct {
	Conditions []Condition
}

func (a *And) Match(i int) bool {
	for _, c := range a.Conditions {
		if !c.Match(i) {
			return false
		}
	}
	return true
}

func NewAnd(conditions ...Condition) *And {
	return &And{
		Conditions: conditions,
	}
}

type Or struct {
	Conditions []Condition
}

func (a *Or) Match(i int) bool {
	for _, c := range a.Conditions {
		if c.Match(i) {
			return true
		}
	}
	return false
}

func NewOr(conditions ...Condition) *Or {
	return &Or{
		Conditions: conditions,
	}
}

type Filter struct {
	data      []int
	condition Condition
}

func (f *Filter) Output() []int {
	var output []int
	for _, v := range f.data {
		if f.condition.Match(v) {
			output = append(output, v)
		}
	}
	return output
}

func main() {
	filter := &Filter{
		data: []int{6, 2, 14, 5, 9, 7, 19, 23, 200},
	}

	// find even and greater than 14
	filter.condition = NewAnd(&Even{}, &Greater{14})
	fmt.Println(filter.Output())

	// find even or greater than 14
	filter.condition = NewOr(&Even{}, &Greater{14})
	fmt.Println(filter.Output())
}
