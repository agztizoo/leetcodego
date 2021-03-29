package easy

import (
    "bytes"
    "strconv"
)

// 38. Count and Say
func CountAndSay(n int) string {
    if n == 1 {
        return "1"
    }
    return explain(CountAndSay(n - 1))
}

func explain(str string) string {
    count, c, res := 1, str[0], bytes.Buffer{}
    for i := 1; i < len(str); i++ {
        if str[i] == c {
            count++
        } else {
            res.WriteString(strconv.Itoa(count))
            res.WriteByte(c)
            count, c = 1, str[i]
        }
    }
    res.WriteString(strconv.Itoa(count))
    res.WriteByte(c)
    return res.String()
}
