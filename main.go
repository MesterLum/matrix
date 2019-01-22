package main

import (
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

	matrix.MulCommutative(&a, &b)
	matrix.MulNonCommutative(&b, &a)
}
