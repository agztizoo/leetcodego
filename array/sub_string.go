package array

func LengthOfLongestSubstring(s string) int {
    if s == "" {
        return 0
    }
    var cache = make(map[string]int)
    var begin = 0
    var max = 0
    for i := 0; i < len(s); i ++ {
        str := s[i:i+1]
        index, exist := cache[str]
        if exist && index >= begin {
            length := i - begin
            if length > max {
                max = length
            }
            begin = index + 1
        }
        cache[str] = i
    }
    len2 := len(s) - begin
    if len2 > max {
        max = len2
    }
    return max
}