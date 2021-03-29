package easy

// 58. Length of Last Word
func LengthOfLastWord(s string) int {
    end := -1
    for i := len(s) - 1; i >= 0; i -- {
        if s[i] != ' ' && end == -1 {
            end = i
        } else if s[i] == ' ' && end != -1 {
            return end - i
        }
    }
    if end == -1 {
        return 0
    }
    return end + 1
}
