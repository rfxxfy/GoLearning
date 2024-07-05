func reverse(x string) string {
	var res string
	for _, elem := range x {
		res = string(elem) + res
	}
	return res
}
