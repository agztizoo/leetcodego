package medium

import "strings"

var (
	romanMap = map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
		4:    "IV",
		9:    "IX",
		40:   "XL",
		90:   "XC",
		400:  "CD",
		900:  "CM",
	}
	romanList = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
)

func intToRoman(num int) string {
	res := strings.Builder{}
	for _, d := range romanList {
		for {
			if num < d {
				break
			}
			num = num - d
			res.WriteString(romanMap[d])
		}
	}
	return res.String()
}
