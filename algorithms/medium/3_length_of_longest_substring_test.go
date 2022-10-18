package medium

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
    tests := []struct {
        name string
        args string
        want int
    }{
        {
            name: "case1",
            args: "abba",
            want: 2,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := lengthOfLongestSubstring(tt.args); got != tt.want {
                t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
            }
        })
    }
}
