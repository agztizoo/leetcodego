package medium

func convert(s string, numRows int) string {
	length, t := len(s), numRows*2-2
	if numRows <= 1 || numRows >= length {
		return s
	}
	arr := make([]byte, 0, length)
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < length; j = j + t {
			arr = append(arr, s[j+i])
			if i > 0 && i < numRows-1 && t+j-i < length {
				arr = append(arr, s[t+j-i])
			}
		}
	}
	return string(arr)
}
