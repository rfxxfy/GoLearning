func check(x string) bool {
	var res string
	for _, elem := range x {
		res = string(elem) + res
	}
	return res == x
}
