package matrix

import (
	"errors"
	"math"
)

// Add function return error or the new matrix with the new values.
func Add(a, b *[][]int) (*[][]int, error) {
	if !ValidateMatrixs(a, b) {
		return nil, errors.New("matrixs no are equals")
	}

	return AddOrSub(a, b, true), nil
}

// Sub function return error or the new matrix with the new values sub.
func Sub(a, b *[][]int) (*[][]int, error) {
	if !ValidateMatrixs(a, b) {
		return nil, errors.New("matrixs no are equals")
	}

	return AddOrSub(a, b, false), nil
}

// Div function divide the matrix.
func Div(a *[][]int, div float64) *[][]float64 {
	var newMatrix [][]float64
	for indexA, sub := range *a {
		var subVector []float64
		for indexB := range sub {
			subVector = append(subVector, float64((*a)[indexA][indexB])/div)
		}
	}
	return &newMatrix
}

// Mul function return the matrix mul.
func Mul(a *[][]int, mul int) *[][]int {
	var newMatrix [][]int
	for indexA, sub := range *a {
		var subVector []int
		for indexB := range sub {
			subVector = append(subVector, (*a)[indexA][indexB]*mul)
		}
	}
	return &newMatrix
}

// MulCommutative ...
func MulCommutative(a, b *[][]int) (*[][]int, error) {
	if !ValidateMatrixsDifferent(a, b) {
		return nil, errors.New("the matrixs are not equals")
	}
	return MultiplicationCon(a, b), nil
}

// MulNonCommutative ...
func MulNonCommutative(a, b *[][]int) (*[][]int, error) {
	if !ValidateMatrixsDifferent(a, b) {
		return nil, errors.New("the matrixs are not equals")
	}
	return MultiplicationCon(a, b), nil
}

// Transponse ...
func Transponse(a *[][]int) *[][]int {
	return InverseMatrix(a)
}

// VectorLength ...
func VectorLength(v *[]int) float64 {
	vLength := 0.0

	for _, v := range *v {
		vLength += math.Pow(float64(v), 2)
	}

	return math.Sqrt(vLength)
}

// DotProduct ...
func DotProduct(a, b *[]int) (int, error) {

	result := 0
	for index := range *a {
		result += (*a)[index] * (*b)[index]
	}

	return result, nil

}
