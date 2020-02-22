package algorithms

import (
    "container/list"
    "math"
)

//7. Reverse Integer
func Reverse(x int) int {
    var rev = 0
    for x != 0 {
        rev = rev * 10 + x % 10;
        x /= 10
        if (rev > math.MaxInt32 || rev < math.MinInt32) {
            return 0
        }
    }
    return rev
}

//9. Palindrome Number
func IsPalindrome2(x int) bool {
    if (x < 0 || (x % 10 == 0 && x != 0)) {
        return false
    }
    var number = 0
    for x > number {
        number = number * 10 + x % 10;
        x = x / 10;
    }
    return x == number || x == number / 10;
}

//13. Roman to Integer
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
        cur := roman(s[i]);
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

//14. Longest Common Prefix
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
        if (len(strs[i]) < minLength) {
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

//20. Valid Parentheses
func ValidParentheses(str string) bool {
    isPrefix := func(b byte) bool {
        if b == '(' || b == '{' || b == '[' {
            return true
        }
        return false
    }
    isPair := func(a, b byte) bool {
        if a == '(' && b == ')' {
            return true
        }
        if a == '{' && b == '}' {
            return true
        }
        if a == '[' && b == ']' {
            return true
        }
        return false
    }
    stack := list.New()
    for i:= 0; i < len(str); i ++ {
        if (isPrefix(str[i])) {
            stack.PushBack(str[i])
        } else {
            if stack.Len() == 0 {
                return false
            }
            bb := stack.Remove(stack.Back())
            if !isPair(bb.(byte), str[i]) {
                return false
            }
        }
    }
    return stack.Len() == 0
}