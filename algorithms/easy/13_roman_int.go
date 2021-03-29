package easy

// 13. Roman to Integer
func RomanToInt(s string) int {
    roman := func(str byte) int {
        switch str {
        case 'I':
            return 1
        case 'V':
            return 5
        case 'X':
            return 10
        case 'L':
            return 50
        case 'C':
            return 100
        case 'D':
            return 500
        case 'M':
            return 1000
        }
        return 0
    }

    count := 0
    for i := 0; i < len(s); i++ {
        cur := roman(s[i])
        next := 0
        if i < len(s)-1 {
            next = roman(s[i+1])
        }
        if cur < next {
            count -= cur
        } else {
            count += cur
        }
    }
    return count
}
