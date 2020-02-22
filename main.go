package main

import (
    "fmt"
    "github.com/leetcodego/array"
    "reflect"
)

type err interface {
    Error() string
}

type RPCError struct {
    Code int64
    Message string
}

func (e *RPCError) Error() string {
   return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
}

//func (e RPCError) Error() string {
//    return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
//}

func NewRPCError(msg string, code int64) *RPCError {
    return &RPCError{
        Code: code,
        Message: msg,
    }
}

func NilOrNot(v interface{}) {
    if v == nil {
        println("nil")
    } else {
        println("Not")
        println(reflect.TypeOf(v).String())
        println(v.(*RPCError))
    }
}

func main() {

    strs := "{}{}{}"

    fmt.Println(array.ValidParentheses(strs))

}

