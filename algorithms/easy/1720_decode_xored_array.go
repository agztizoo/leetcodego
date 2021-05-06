package easy

func decode(encoded []int, first int) []int {
	if len(encoded) == 0 {
		return []int{}
	}
	arr := make([]int, len(encoded)+1)
	arr[0] = first
	for i, v := range encoded {
		arr[i+1] = v ^ arr[i]
	}
	return arr
}
