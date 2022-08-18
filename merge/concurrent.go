package merge

func Concurrent(a []int, res chan []int) {
	if len(a) <= 1 {
		res <- a
		return
	}

	middleLen := len(a) / 2
	resRight := make(chan []int, middleLen)
	defer close(resRight)
	resLeft := make(chan []int, middleLen)
	defer close(resLeft)

	go Concurrent(a[:middleLen], resRight)
	go Concurrent(a[middleLen:], resLeft)

	left := <-resLeft
	right := <-resRight

	if left[len(left)-1] <= right[0] {
		l := make([]int, len(left))
		copy(l, left)
		res <- append(l, right...)
		return
	}

	res <- mergeConcurrent(right, left)
}

func mergeConcurrent(right, left []int) []int {
	merged := make([]int, 0, len(right)+len(left))

	for len(right) > 0 && len(left) > 0 {
		if right[0] >= left[0] {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}

	if len(right) > 0 {
		merged = append(merged, right...)
	}

	if len(left) > 0 {
		merged = append(merged, left...)
	}

	return merged
}
