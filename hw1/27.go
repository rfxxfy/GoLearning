func mergeSortedArrays(arr1 []int, arr2 []int) []int {
	len1, len2 := len(arr1), len(arr2)
	res := make([]int, len1+len2)

	originalArr1 := make([]int, len1)
	copy(originalArr1, arr1)
	originalArr2 := make([]int, len2)
	copy(originalArr2, arr2)

	arr1 = append(arr1, math.MaxInt64)
	arr2 = append(arr2, math.MaxInt64)

	i, j := 0, 0
	for k := 0; k < len1+len2; k++ {
		if arr1[i] < arr2[j] {
			res[k] = arr1[i]
			i++
		} else {
			res[k] = arr2[j]
			j++
		}
	}

	copy(arr1, originalArr1)
	copy(arr2, originalArr2)

	return res
}

/*
func mergeSortedArrays(arr1 []int, arr2 []int) []int {
  merged := append(arr1Copy, arr2Copy...)
	sort.Ints(merged)
	return merged
}
*/
