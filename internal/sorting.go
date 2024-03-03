package internal

// Sort the given numbers and return them as new sorted numbers while preserving the original numbers.
func BubbleSort(originalNums []int) (newNums []int) {
	newNums = make([]int, len(originalNums))
	copy(newNums, originalNums)

	swapping := true
	for end := len(newNums) - 1; swapping; {
		swapping = false

		for i := 0; i < end; i++ {
			if newNums[i] > newNums[i+1] {
				newNums[i], newNums[i+1] = newNums[i+1], newNums[i]
				swapping = true
			}
		}

		end -= 1
	}

	return newNums
}
