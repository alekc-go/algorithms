package selection

func Sort(data []int) []int {
	var currentValue, minValue, minValuePosition int
	for i := 0; i < len(data); i++ {
		currentValue = data[i]

		//find minimum value
		minValue = data[i]
		minValuePosition = i
		for j := i + 1; j < len(data); j++ {
			if data[j] < minValue {
				minValue = data[j]
				minValuePosition = j
			}
		}
		data[i] = minValue
		data[minValuePosition] = currentValue
	}
	return data
}
