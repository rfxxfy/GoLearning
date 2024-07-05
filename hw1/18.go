func minAndMax(mass []int) (int, int) {
	min_ := mass[0]
	max_ := mass[0]
	for _, m := range mass {
		if m < min_ {
			min_ = m
		}
		if m > max_ {
			max_ = m
		}
	}
}
