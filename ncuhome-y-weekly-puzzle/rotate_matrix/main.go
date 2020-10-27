package main

import "fmt"

func index_rotater(size int) func(x, y int) (int, int) {
	return func(x, y int) (int, int) {
		return y, size - 1 - x
	}
}

func new_rotate(matrix [][]int) [][]int {
	size := len(matrix)
	new := make([][]int, size)
	rotater := index_rotater(size)

	for i := range matrix {
		new[i] = make([]int, size)
	}

	for x1 := range matrix {
		for y1 := range matrix {
			x2, y2 := rotater(x1, y1)
			new[x2][y2] = matrix[x1][y1]
		}
	}

	return new
}

func inplace_rotate(matrix [][]int) [][]int {
	size := len(matrix)
	rotater := index_rotater(size)

	for x1 := 0; x1 < size/2; x1++ {
		for y1 := 0; y1 < (size+1)/2; y1++ {
			x2, y2 := rotater(x1, y1)
			x3, y3 := rotater(x2, y2)
			x4, y4 := rotater(x3, y3)

			matrix[x2][y2], matrix[x3][y3], matrix[x4][y4], matrix[x1][y1] = matrix[x1][y1], matrix[x2][y2], matrix[x3][y3], matrix[x4][y4]
		}
	}

	return matrix
}

func main() {
	testcases := [][][]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		},
	}

	for _, testcase := range testcases {
		fmt.Println("- Results -")

		result1 := new_rotate(testcase)
		result2 := inplace_rotate(testcase)

		for i := range testcase {
			fmt.Println(result1[i], result2[i])
		}
	}
}
