package piscine

func AppendRange(min, max int) []int {
	tab := []int{}
	for min < max {
		tab = append(tab, min)
		min++
	}
	return tab
}
