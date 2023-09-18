package piscine

func ConcatParams(args []string) string {
	r := ""
	i := 0
	for range args {
		r += string(args[i])
		if i != len(args)-1 {
			r += "\n"
		}
		i++
	}
	return (r)
}
