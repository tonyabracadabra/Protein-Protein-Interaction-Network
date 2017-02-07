package main

import (
	"fmt"
	"os"
	"strconv"
)

// vertex name must be nonnegative integers
type Edge struct {
	v int
	w int
	weight float64
}

func CreateNewEdge(v, w int, weight float64) Edge {
	if v < 0 || w < 0{
		fmt.Println("Vertex name must be nonnegative integers!")
		os.Exit(1)
	}

	if weight < 0 {
		fmt.Println("Weight must be larger than 0!")
		os.Exit(1)
	}

	return Edge{v, w, weight}
}

func (e *Edge) From() int {
	return e.v
}

func (e *Edge) To() int {
	return e.w
}

func (e *Edge) Weight() float64 {
	return e.weight
}

func (e *Edge) ToString() string {
	return strconv.Itoa(e.v) + "->" + strconv.Itoa(e.w) + " " + strconv.FormatFloat(e.weight, 'f', 2, 64)
}

func (e *Edge) PrintEdge() {
	fmt.Println(e.ToString())
}