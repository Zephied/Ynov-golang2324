package piscine

func Concat(str1 string, str2 string) string {
	i := 0
	r := ""
	for range str1 {
		r += string(str1[i])
		i++
	}
	j := 0
	for range str2 {
		r += string(str2[j])
		j++
	}
	return (r)
}
