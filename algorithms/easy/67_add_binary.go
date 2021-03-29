package easy

import "bytes"

// 67. Add Binary
func AddBinary(a string, b string) string {
    if a == "" || a == "0" {
        return b
    }
    if b == "" || b == "0" {
        return a
    }
    temp, buff := 0, bytes.Buffer{}
    for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; {
        if i >= 0 && a[i] == '1' {
            temp += 1
        }
        if j >= 0 && b[j] == '1' {
            temp += 1
        }
        if temp == 0 || temp == 2 {
            buff.WriteString("0")
        } else {
            buff.WriteString("1")
        }
        if temp < 2 {
            temp = 0
        } else {
            temp = 1
        }
        i--
        j--
    }
    if temp != 0 {
        buff.WriteString("1")
    }
    for i, j := 0, len(buff.Bytes())-1; i < j; {
        buff.Bytes()[i], buff.Bytes()[j] = buff.Bytes()[j], buff.Bytes()[i]
        i++
        j--

    }
    return buff.String()
}
