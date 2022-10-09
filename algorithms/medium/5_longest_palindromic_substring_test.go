package medium

import "testing"

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "case1",
			args: "babad",
			want: "aba",
		},
		{
			name: "case2",
			args: "aabbcc",
			want: "bb",
		},
		{
			name: "case3",
			args: "a",
			want: "a",
		},
		{
			name: "case4",
			args: "ab",
			want: "b",
		},
		{
			name: "case5",
			args: "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
