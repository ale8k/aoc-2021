package aoc1

func AOC1P1(data []int) int {
	increasedCount := 0
	for i := 1; i < len(data); i++ {
		if data[i-1] < data[i] {
			increasedCount++
		}
	}
	return increasedCount
}

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
