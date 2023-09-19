package piscine

func Map(f func(int) bool, a []int) []bool {
	tab := make([]bool, 0)
	for _, v := range a {
		tab = append(tab, f(v))
	}
	return tab
}
