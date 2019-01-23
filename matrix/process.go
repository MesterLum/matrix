package matrix

import (
	"errors"
	"math"
)

// Add function return error or the new matrix with the new values.
func Add(a, b *[][]float64) (*[][]float64, error) {
	if !ValidateMatrixs(a, b) {
		return nil, errors.New("matrixs no are equals")
	}

	return AddOrSub(a, b, true), nil
}

// Sub function return error or the new matrix with the new values sub.
func Sub(a, b *[][]float64) (*[][]float64, error) {
	if !ValidateMatrixs(a, b) {
		return nil, errors.New("matrixs no are equals")
	}

	return AddOrSub(a, b, false), nil
}

// Div function divide the matrix.
func Div(a *[][]float64, div float64) *[][]float64 {
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
func Mul(a *[][]float64, mul float64) *[][]float64 {
	var newMatrix [][]float64
	for indexA, sub := range *a {
		var subVector []float64
		for indexB := range sub {
			subVector = append(subVector, (*a)[indexA][indexB]*mul)
		}
	}
	return &newMatrix
}

// MulCommutative ...
func MulCommutative(a, b *[][]float64) (*[][]float64, error) {
	if !ValidateMatrixsDifferent(a, b) {
		return nil, errors.New("the matrixs are not equals")
	}
	return MultiplicationCon(a, b), nil
}

// MulNonCommutative ...
func MulNonCommutative(a, b *[][]float64) (*[][]float64, error) {
	/*if !ValidateMatrixsDifferent(a, b) {
		return nil, errors.New("the matrixs are not equals")
	}*/
	return MultiplicationCon(a, b), nil
}

// Transponse ...
func Transponse(a *[][]float64) *[][]float64 {
	return InverseMatrix(a)
}

// VectorLength ...
func VectorLength(v *[]float64) float64 {
	vLength := 0.0

	for _, v := range *v {
		vLength += math.Pow(float64(v), 2)
	}

	return math.Sqrt(vLength)
}

// DotProduct ...
func DotProduct(a, b *[]float64) (float64, error) {

	var result float64 = 0.0
	for index := range *a {
		result += (*a)[index] * (*b)[index]
	}

	return result, nil

}

// BooleanToBipolar ...
func BooleanToBipolar(a *[][]float64) *[][]float64 {
	var newMatrix [][]float64

	for _, sub := range *a {
		var subVector []float64
		for _, value := range sub {
			subVector = append(subVector, float64(2*value-1))
		}
		newMatrix = append(newMatrix, subVector)
	}
	return &newMatrix
}

// BipolarToBoolean ...
func BipolarToBoolean(a *[][]float64) *[][]float64 {
	var newMatrix [][]float64

	for _, sub := range *a {
		var subVector []float64
		for _, value := range sub {
			subVector = append(subVector, float64((value+1)/2))
		}
		newMatrix = append(newMatrix, subVector)
	}
	return &newMatrix
}
