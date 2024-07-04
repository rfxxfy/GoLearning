func reverse(x string) string {
	a := []rune(x)
	var res string
	for _, elem := range a {
		res = string(elem) + res
	}
	return res
}
