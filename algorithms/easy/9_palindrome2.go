package easy

// 9. Palindrome Number
func IsPalindrome2(x int) bool {
    if x < 0 || (x%10 == 0 && x != 0) {
        return false
    }
    numer := 0
    for x > numer {
        numer = numer*10 + x%10
        x = x / 10
    }
    return x == numer || x == numer/10
}
