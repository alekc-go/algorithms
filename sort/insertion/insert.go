package insertion

func Sort(data []int) []int {
	//it's already sorted if we have less than 2 elements
	if len(data) < 2 {
		return data
	}

	//first element is already sorted, let's start with the second one
	var j int
	for i := 1; i < len(data); i++ {
		value := data[i]

		//if the current value is bigger than the previous entry, then shift the array
		for j = i; j > 0 && value < data[j-1]; j-- {
			data[j] = data[j-1]
		}
		//now that our array is completely shifted, assign the value where we stopped
		data[j] = value
	}
	return data
}

func SortRecursive(data []int, i int, n int) []int {
	value := data[i]
	var j int
	for j = i; j > 0 && value < data[j-1]; j-- {
		data[j] = data[j-1]
	}
	data[j] = value
	if i+1 <= n {
		return SortRecursive(data, i+1, n)
	}
	return data
}
