package matrix

// ValidateMatrixs check if both matrix are equals.
func ValidateMatrixs(a, b *[][]float64) bool {
	return len(*a) == len(*b) && len((*a)[0]) == len((*b)[0])
}

// ValidateMatrixsDifferent check if both matrix are equals or equivalents
func ValidateMatrixsDifferent(a, b *[][]float64) bool {
	return len(*a) == len((*b)[0]) && len(*b) == len((*a)[0])
}

// AddOrSub is a function for provide help and minify the source
func AddOrSub(a, b *[][]float64, isAdd bool) *[][]float64 {
	var newMatrix [][]float64
	for indexA, sub := range *a {
		var subVector []float64
		for indexB := range sub {
			if isAdd {
				subVector = append(subVector, (*a)[indexA][indexB]+(*b)[indexA][indexB])
			} else {
				subVector = append(subVector, (*a)[indexA][indexB]-(*b)[indexA][indexB])
			}

		}
		newMatrix = append(newMatrix, subVector)
	}
	return &newMatrix
}

// InverseMatrix ...
func InverseMatrix(a *[][]float64) *[][]float64 {
	newMatrix := make([][]float64, len((*a)[0]))
	for i := range newMatrix {
		newMatrix[i] = make([]float64, len(*a))
	}
	for indexA, sub := range *a {

		for indexB := range sub {
			newMatrix[indexB][indexA] = (*a)[indexA][indexB]

		}

	}
	return &newMatrix
}

// MultiplicationCon ...
func MultiplicationCon(a, b *[][]float64) *[][]float64 {
	var newMatrix [][]float64

	for _, subA := range *a {
		var subVector []float64
		for indexBB := range (*b)[0] {
			var accum float64 = 0.0
			for indexAA, d := range subA {
				accum += d * (*b)[indexAA][indexBB]
			}
			subVector = append(subVector, accum)
		}
		newMatrix = append(newMatrix, subVector)

	}

	return &newMatrix
}
