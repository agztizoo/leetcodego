package easy

import "container/list"

// 20. Valid Parentheses
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
    for i := 0; i < len(str); i++ {
        if isPrefix(str[i]) {
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
