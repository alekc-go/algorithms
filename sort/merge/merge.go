package merge

func Sort(data []int) []int {
	length := len(data)
	middle := length / 2
	if length < 2 {
		return data
	}
	//Continue recursively to split all elements until we reach the point where we only have
	// 2 arrays with <=1 element in it
	return mergeSlice(Sort(data[:middle]), Sort(data[middle:]))
}

func mergeSlice(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	// Merge 2 slices together by popping the smallest element between
	// two first elements (1 in each slice).
	// once exausted all elements in one of the slices just add remaining to the end of a queue.
	// there will never be more than n+1 elmemnts in one slice compared to another
	i := 0
	//if both left and right are not empty
	for ; len(left) > 0 && len(right) > 0; i++ {
		//if initial element of left array is smaller than the initial element of right array
		if left[0] < right[0] {
			//assign first element of let to the result in curr position, and pop first element from the left
			result[i] = left[0]
			left = left[1:]
			continue
		}
		//since the right is bigger, pop the first element and put it in result
		result[i] = right[0]
		right = right[1:]
	}

	//for any left over from left array put it at the end of the result
	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	//for any left over from right array, append it to the result
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
