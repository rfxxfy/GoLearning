func check(x string) bool {
	var res string
	for _, elem := range a {
		res = string(elem) + res
	}
	return res == x
}
