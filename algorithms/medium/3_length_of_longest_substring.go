package medium

func lengthOfLongestSubstring(s string) int {
    if s == "" {
        return 0
    }
    var cache = make(map[int32]int)
    begin, max := 0, 0
    for i, x := range s {
        if index, exist := cache[x]; exist && index >= begin {
            if i-begin > max {
                max = i - begin
            }
            begin = index + 1
        }
        cache[x] = i
    }
    len2 := len(s) - begin
    if len2 > max {
        max = len2
    }
    return max
}
