package easy

// 9. Palindrome Number
func IsPalindrome2(x int) bool {
    if x < 0 || (x%10 == 0 && x != 0) {
        return false
    }
    var number = 0
    for x > number {
        number = number*10 + x%10
        x = x / 10
    }
    return x == number || x == number/10
}
