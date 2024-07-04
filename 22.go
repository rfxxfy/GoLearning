func bubbleSort(mass []int) {
	len_ := len(mass)
	for i := 0; i < len_-1; i++ {
		for j := 0; j < len_-i-1; j++ {
			if mass[j] > mass[j+1] {
				mass[j], mass[j+1] = mass[j+1], mass[j]
			}
		}
	}
}
