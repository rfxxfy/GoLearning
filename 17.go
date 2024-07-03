func check(x string) bool {
	var res string
	for i := len(x) - 1; i >= 0; i-- {
		res += string(x[i])
	}
	return res == x
}
