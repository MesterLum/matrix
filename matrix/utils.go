package matrix

// ValidateMatrixs check if both matrix are equals.
func ValidateMatrixs(a, b *[][]int) bool {
	return true // Temporary
}

// AddOrSub is a function for provide help and minify the source
func AddOrSub(a, b *[][]int, isAdd bool) *[][]int {
	var newMatrix [][]int
	for indexA, sub := range *a {
		var subVector []int
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
func InverseMatrix(a *[][]int) *[][]int {
	newMatrix := make([][]int, len((*a)[0]))
	for i := range newMatrix {
		newMatrix[i] = make([]int, len(*a))
	}
	for indexA, sub := range *a {

		for indexB := range sub {
			newMatrix[indexB][indexA] = (*a)[indexA][indexB]

		}

	}
	return &newMatrix
}

// MultiplicationCon ...
func MultiplicationCon(a, b *[][]int) *[][]int {
	var newMatrix [][]int

	for _, subA := range *a {
		var subVector []int
		for indexBB := range (*b)[0] {
			accum := 0
			for indexAA, d := range subA {
				accum += d * (*b)[indexAA][indexBB]
			}
			subVector = append(subVector, accum)
			accum = 0
		}
		newMatrix = append(newMatrix, subVector)

	}

	return &newMatrix
}
