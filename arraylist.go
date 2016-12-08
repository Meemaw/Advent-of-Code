package advent_of_code

import (
	"reflect"
	"fmt"
)

type List struct {
	elements 		[]interface{}
	num_elements	int
}

// Insantiates a new list
func NewList() *List {
	return &List{}
}

const (
	growth_factor = float32(2.0) // double capcity of the list
	shrink_factor = float32(0.25) // shrink when num_elements reaches 25% of capacity
)

func(list *List) Elements() []interface{} {
	return list.elements
}

func(list *List) resize(new_capacity int) {
	new_elements := make([]interface{}, new_capacity)
	copy(new_elements, list.elements)
	list.elements = new_elements
}

// Adds values to the end of the list
func(list *List) Add(elements ...interface{}) {
	list.growBy(len(elements))
	for _, element := range elements {
		list.elements[list.num_elements] = element
		list.num_elements++
	}
}

func(list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}

	return list.elements[index], true
}

func(list *List) Contains(elements ...interface{}) bool {
	for _, element := range elements {
		for i := 0; i < list.num_elements; i++ {
			if(reflect.DeepEqual(element, list.elements[i])) {
				return true
			}
		}
	}
	return false
}

// Expands the list by growth_factor if necesary, i.e num_elements reaches capacity
func(list* List) growBy(n int) {
	current_capacity := cap(list.elements)
	if list.num_elements + n >= current_capacity {
		new_capacity := int(growth_factor * float32(current_capacity + n))
		list.resize(new_capacity)
	}
}

func(list* List) Print() {
	for i := 0; i < list.num_elements; i++ {
		fmt.Print(list.elements[i])
		fmt.Print("\n")
	}
}

func(list* List) isEmpty() bool {
	return list.num_elements == 0
}

func(list* List) Size() int {
	return list.num_elements
}

// Checks whether index is within list bounds
func(list *List) withinRange(index int) bool {
	return index >= 0 && index < list.num_elements
}

// Remove all elements from the list
func(list *List) clear() {
	list.num_elements = 0
	list.elements = []interface{}{}
}

// Remove element at given index
func(list* List) remove(index int) {
	if(!list.withinRange(index)) {
		return
	}

	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.num_elements])
	list.num_elements--
	list.shrink()
}

// Shrinks the list if necesary, i.e num_elements reaches shrink_factor
func(list* List) shrink() {
	current_capacity := cap(list.elements)
	if list.num_elements <= int(float32(current_capacity) * shrink_factor) {
		list.resize(list.num_elements)
	}
}
