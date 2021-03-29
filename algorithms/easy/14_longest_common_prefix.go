package easy

import "math"

// 14. Longest Common Prefix
func LongestCommonPrefix(strs []string) string {
    isCommonPrefix := func(strs []string, length int) bool {
        var str = strs[0][0:length]
        for i := 1; i < len(strs); i++ {
            if !(strs[i][0:length] == str) {
                return false
            }
        }
        return true
    }
    if strs == nil || len(strs) == 0 {
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }
    var minLength = math.MaxInt32
    for i := 0; i < len(strs); i++ {
        if len(strs[i]) < minLength {
            minLength = len(strs[i])
        }
    }
    var index = 1
    for index <= minLength {
        middle := (index + minLength) / 2
        if isCommonPrefix(strs, middle) {
            index = middle + 1
        } else {
            minLength = middle - 1
        }
    }
    return strs[0][0 : (index+minLength)/2]
}
