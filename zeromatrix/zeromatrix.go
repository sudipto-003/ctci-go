package zeromatrix

type Matrix [][]int

func ConvertZeroMatrix(a Matrix) Matrix {
	aM := len(a)
	if aM == 0 {
		return a
	}
	aN := len(a[0])

	rows := make([]bool, aM)
	coloums := make([]bool, aN)

	for i := 0; i < aM; i++ {
		for j := 0; j < aN; j++ {
			if a[i][j] == 0 {
				rows[i] = true
				coloums[j] = true
			}
		}
	}

	for i := 0; i < len(rows); i++ {
		if rows[i] {
			makeRowZero(a, i)
		}
	}

	for j := 0; j < len(coloums); j++ {
		if coloums[j] {
			makeColoumZero(a, j)
		}
	}

	return a
}

func makeRowZero(a Matrix, i int) {
	n := len(a[0])

	for k := 0; k < n; k++ {
		a[i][k] = 0
	}
}

func makeColoumZero(a Matrix, j int) {
	m := len(a)

	for k := 0; k < m; k++ {
		a[k][j] = 0
	}
}
