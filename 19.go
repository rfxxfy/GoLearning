func remove(x []int, ind int) {
	x = append(x[:ind], x[ind+1:]...)
}
