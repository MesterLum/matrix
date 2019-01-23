package main

import (
	"fmt"

	"./matrix"
)

func main() {
	var ma1 = [][]float64{
		{3, -1, 5, 4},
		{0, 7, 8, -3},
		{2, 9, 11, 6},
		{-4, 2, 5, 0},
	}

	var m2 = [][]float64{
		{0.09, 0.0862, -0.015, -0.192},
		{-0.17, -0.003, 0.1104, -0.071},
		{0.14, 0.0701, -0.056, 0.0744},
		{-0.03, -0.153, 0.1085, 0.0336},
	}

	var b = [][]float64{
		{6},
		{0},
		{11},
		{-18},
	}

	if r, err := matrix.MulNonCommutative(&ma1, &m2); err != nil {
		panic(err)
	} else {
		for _, sub := range *r {
			for _, value := range sub {
				fmt.Printf("%0.1f ", value)
			}
			fmt.Println()
		}
	}
	fmt.Println()
	//fmt.Println(m2)

	if r, err := matrix.MulNonCommutative(&m2, &b); err != nil {
		panic(err)
	} else {
		var subVector []float64
		for _, sub := range *r {

			for _, value := range sub {
				subVector = append(subVector, value)

			}

		}
		fmt.Println(subVector)
		result := matrix.Transponse(&ma1)

		for _, sub := range *result {
			fmt.Println()
			resuli, _ := matrix.DotProduct(&sub, &subVector)
			fmt.Println(resuli)

		}

	}

}
