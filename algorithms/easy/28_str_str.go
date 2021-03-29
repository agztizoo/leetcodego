package easy

// 28. Implement strStr()
func StrStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    pLength := len(needle)
    txtLength := len(haystack)
    if haystack == "" || pLength > txtLength {
        return -1
    }
    idx := 0
    for idx+pLength <= txtLength {
        if haystack[idx] == needle[0] && haystack[idx:idx+pLength] == needle {
            return idx
        }
        idx++
    }
    return -1
}
