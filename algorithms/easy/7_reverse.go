package easy

import "math"

// 7. Reverse Integer
func Reverse(x int) int {
    var rev = 0
    for x != 0 {
        rev = rev*10 + x%10
        x /= 10
        if rev > math.MaxInt32 || rev < math.MinInt32 {
            return 0
        }
    }
    return rev
}
