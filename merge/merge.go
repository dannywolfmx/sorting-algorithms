package merge

func Mergesort(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	right := Mergesort(a[0 : len(a)/2])
	left := Mergesort(a[len(a)/2:])

	if left[len(left)-1] <= right[0] {
		l := make([]int, len(left))
		copy(l, left)
		return append(l, right...)
	}

	return merge(right, left)
}

func merge(firstSlice, secondSlice []int) []int {
	merged := make([]int, 0, len(firstSlice)+len(secondSlice))

	for len(firstSlice) > 0 && len(secondSlice) > 0 {
		if firstSlice[0] >= secondSlice[0] {
			merged = append(merged, secondSlice[0])
			secondSlice = secondSlice[1:]
		} else {
			merged = append(merged, firstSlice[0])
			firstSlice = firstSlice[1:]
		}
	}

	if len(firstSlice) > 0 {
		merged = append(merged, firstSlice...)
	}

	if len(secondSlice) > 0 {
		merged = append(merged, secondSlice...)
	}

	return merged
}
