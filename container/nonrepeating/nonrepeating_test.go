package main

import (
	"testing"
)

// 测试 - 代码覆盖率 testing.T
func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcdddd", 4},
		{"lllllll", 1},
		{"", 0},
		{"我离歌曲", 4},
		{"吧里bs防化服好", 8},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d",
				actual, tt.s, tt.ans)
		}
	}
}

// 性能测试 testing.B
func BenchmarkSubstr(b *testing.B) {
	s := "吧里bs防化服好"
	for i := 0; i < 2; i++ {
		s = s +s
	}

	b.Logf("len(s) = %d", len(s))

	// 上线拼接字符串的时间不算, 重置计时
	b.ResetTimer()

	ans := 8

	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d",
				actual, s, ans)
		}
	}

}
