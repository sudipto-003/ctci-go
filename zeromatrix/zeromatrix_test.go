package zeromatrix

import "testing"

func testMatrixEqual(t *testing.T, test int, a, b Matrix) bool {
	if len(a) == 0 || len(b) == 0 {
		return true
	}
	aM, aN := len(a), len(a[0])
	bM, bN := len(b), len(b[0])

	if aM != bM || aN != bN {
		t.Errorf("Error in test: %d. Matrixes dimension does not match", test)
		return false
	}

	for i := 0; i < aM; i++ {
		for j := 0; j < aN; j++ {
			if a[i][j] != b[i][j] {
				t.Errorf("Error in test: %d. Elements does not match. want=%d got=%d", test, a[i][j], b[i][j])
				return false
			}
		}
	}
	return true
}

func TestZeroMatrix(t *testing.T) {
	tests := []struct {
		original  Matrix
		converted Matrix
	}{
		{
			Matrix{[]int{1, 2}, []int{3, 0}},
			Matrix{[]int{1, 0}, []int{0, 0}},
		},
		{
			Matrix{[]int{1, 2}, []int{3, 4}},
			Matrix{[]int{1, 2}, []int{3, 4}},
		},
		{
			Matrix{[]int{1}},
			Matrix{[]int{1}},
		},
		{
			Matrix{[]int{1, 0, 3}, []int{4, 5, 6}, []int{7, 8, 9}},
			Matrix{[]int{0, 0, 0}, []int{4, 0, 6}, []int{7, 0, 9}},
		},
		{
			Matrix{[]int{1, 0, 3}, []int{0, 5, 6}, []int{7, 8, 9}},
			Matrix{[]int{0, 0, 0}, []int{0, 0, 0}, []int{0, 0, 9}},
		},
		{
			Matrix{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}, []int{5, 5, 5}},
			Matrix{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}, []int{5, 5, 5}},
		},
		{
			Matrix{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 0, 9}, []int{5, 5, 5}},
			Matrix{[]int{1, 0, 3}, []int{4, 0, 6}, []int{0, 0, 0}, []int{5, 0, 5}},
		},
		{
			Matrix{[]int{1, 2, 0, 7}, []int{4, 5, 6, 8}, []int{7, 8, 9, 9}, []int{5, 5, 5, 6}, []int{3, 6, 9, 1}},
			Matrix{[]int{0, 0, 0, 0}, []int{4, 5, 0, 8}, []int{7, 8, 0, 9}, []int{5, 5, 0, 6}, []int{3, 6, 0, 1}},
		},
		{
			Matrix{[]int{}},
			Matrix{[]int{}},
		},
		{
			Matrix{},
			Matrix{},
		},
		{
			Matrix{[]int{1}},
			Matrix{[]int{1}},
		},
	}

	for i, tt := range tests {
		converted := ConvertZeroMatrix(tt.original)
		testMatrixEqual(t, i, converted, tt.converted)
	}
}
