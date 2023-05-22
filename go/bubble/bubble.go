package bubble

func Sort(array []int) []int {
	for {
		changed := false
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				temp := array[i]
				array[i] = array[i+1]
				array[i+1] = temp
				changed = true
			}
		}
		if !changed {
			break
		}
	}
	return array
}

func SortRecur(array []int, length int) {
	for i := 0; i < length-1; i++ {
		if array[i] > array[i+1] {
			temp := array[i]
			array[i] = array[i+1]
			array[i+1] = temp
		}
	}
	if length == 0 {
		return
	}
	SortRecur(array, length-1)
}
