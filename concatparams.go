package piscine

func ConcatParams(args []string) string {
	r := ""
	i := 0
	for range args {
		r += string(args[i])
		r += "\n"
		i++
	}
	return (r)
}
