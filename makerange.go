package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return []int(nil)
	}
	tab := make([]int, max-min)
	i := 0
	for min < max {
		tab[i] = min
		i++
		min++
	}
	return tab
}
