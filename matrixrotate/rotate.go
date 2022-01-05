package matrixrotate

func Rotate(matrix [][]int32) bool {
	if len(matrix) == 0 || (len(matrix[0]) != len(matrix)) {
		return false
	}

	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first

			top := matrix[first][i]

			matrix[first][i] = matrix[last-offset][first]

			matrix[last-offset][first] = matrix[last][last-offset]

			matrix[last][last-offset] = matrix[i][last]

			matrix[i][last] = top
		}
	}

	return true
}
