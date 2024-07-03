func reverse(x string) string {
	var res string
	for i := len(x) - 1; i >= 0; i-- {
		res += string(x[i])
	}
	return res
}
