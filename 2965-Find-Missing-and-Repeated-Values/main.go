package main

func findMissingAndRepeatedValues(grid [][]int) []int {
	ans := make([]int, 2)
	n := len(grid)
	numbers := make([]int, n*n)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			numbers[grid[r][c]-1]++
			if numbers[grid[r][c]-1] == 2 {
				ans[0] = grid[r][c]
			}
		}
	}

	for i, n := range numbers {
		if n == 0 {
			ans[1] = i + 1
		}
	}

	return ans

}
