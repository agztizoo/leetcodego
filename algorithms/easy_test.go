package algorithms

import (
    "fmt"
    "testing"

    "github.com/leetcodego/algorithms/easy"
)

func TestSearchInsert(t *testing.T) {
    nums := []int{1,4,6,7,8,9}
    target := 7
    idx := easy.SearchInsert(nums, target)
    if idx != 3 {
        t.Logf("error, input=%d", target)
        t.FailNow()
    }
    target = 0
    idx = easy.SearchInsert(nums, target)
    if idx != 0 {
        t.Logf("error, input=%d", target)
        t.FailNow()
    }
    target = 10
    idx = easy.SearchInsert(nums, target)
    if idx != 6 {
        t.Logf("error, input=%d", target)
        t.FailNow()
    }
    target = 2
    idx = easy.SearchInsert(nums, target)
    if idx != 1 {
        t.Logf("error, input=%d", target)
        t.FailNow()
    }
}

func TestCountAndSay(t *testing.T) {
    level := 5
    res := easy.CountAndSay(level)
    if res != "111221" {
        t.Logf("error, input=%d", level)
        t.FailNow()
    }
}

func TestLengthOfLastWord(t *testing.T) {
    a, b := "11110", "101"
    fmt.Println(easy.AddBinary(a, b))
}
