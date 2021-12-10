package aoc1

// AOC Day 1 challenge part 1
func AOC1P1(data []int) int {
	increasedCount := 0
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			increasedCount++
		}
	}
	return increasedCount
}
