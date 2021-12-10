package aoc1

// AOC Day 1 challenge part 2
func AOC1P2(data []int) int {
	increasedWindowCount := 0
	for i := range data {
		if i+2 > len(data)-1 || i+3 > len(data)-1 {
			return increasedWindowCount
		}
		firstWindow := data[i] + data[i+1] + data[i+2]
		secondWindow := data[i+1] + data[i+2] + data[i+3]
		if secondWindow > firstWindow {
			increasedWindowCount++
		}
	}
	return increasedWindowCount
}
