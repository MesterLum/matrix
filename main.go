package main

import (
	"fmt"

	"./matrix"
)

func main() {
	var a = [][]int{
		{1, 4},
		{2, 5},
		{3, 6},
		{4, 7},
	}
	var b = [][]int{
		{7, 8, 9},
		{10, 11, 12},
		{10, 12, 15},
	}
	var a1 = [][]int{
		{1, 2},
		{2, 3},
	}

	var b1 = [][]int{
		{1, 2},
		{2, 3},
	}

	var a2 = [][]int{
		{1, 2},
		{3, 4},
		{5, 7},
	}
	var b2 = [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	if _, err := matrix.Add(&a1, &b1); err != nil {
		panic(err)
	}

	if r, err := matrix.MulCommutative(&a2, &b2); err != nil {
		panic(err)
	} else {
		fmt.Println(r)
	}

	matrix.MulCommutative(&a, &b)
	matrix.MulNonCommutative(&b, &a)
}
