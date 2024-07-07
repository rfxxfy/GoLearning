type Rectangle struct {
	height int
	width  int
}

func (x *Rectangle) area() int {
	return x.width * x.height
}
