package piscine

func AppendRange(min, max int) []int {
	tab := []int{}
	if min >= max {
		return []int(nil)
	}
	for min < max {
		tab = append(tab, min)
		min++
	}
	return tab
}
