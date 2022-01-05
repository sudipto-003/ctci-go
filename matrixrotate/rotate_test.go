package matrixrotate

import (
	"testing"
)

type test struct {
	matrix  [][]int32
	rotated [][]int32
}

func matrixequal(mat1, mat2 [][]int32) bool {
	if (len(mat1) != len(mat2)) || (len(mat1[0]) != len(mat2[0])) {
		return false
	}

	for i := 0; i < len(mat1); i++ {
		for j := 0; j < len(mat1[0]); j++ {
			if mat1[i][j] != mat2[i][j] {
				return false
			}
		}
	}

	return true
}

var testcases = []test{
	{
		[][]int32{[]int32{5, 6}, []int32{6, 4}},
		[][]int32{[]int32{6, 5}, []int32{4, 6}},
	},
	{
		[][]int32{[]int32{1, 2, 7}, []int32{3, 4, 8}, []int32{5, 6, 9}},
		[][]int32{[]int32{5, 3, 1}, []int32{6, 4, 2}, []int32{9, 8, 7}},
	},
	{
		[][]int32{[]int32{5}},
		[][]int32{[]int32{5}},
	},
}

func TestMatrix(t *testing.T) {
	for _, test := range testcases {
		Rotate(test.matrix)
		if !matrixequal(test.matrix, test.rotated) {
			t.Errorf("Error on matrix rotate")
		}
	}
}
