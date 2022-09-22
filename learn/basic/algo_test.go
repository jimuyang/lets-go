package main

import (
	"math"
	"testing"
)

// go test . 执行测试
// go test -coverprofile=c.out . 覆盖测试
// go tool cover -html=c.out 查看覆盖测试结果
func TestAdd(t *testing.T) {
	cases := []struct{ a, b, c int }{
		{1, 2, 3},
		{1, -1, 0},
		{math.MaxInt64, 1, math.MinInt64},
		// {1, 2, 1},
	}
	for _, c := range cases {
		if res := Add(c.a, c.b); res != c.c {
			t.Errorf("add(%d, %d) result: %d, expect: %d", c.a, c.b, res, c.c)
		}
	}
}

// go test -bench .
func BenchmarkAdd(b *testing.B) {
	a1, a2, c := 1, 2, 3
	for i := 0; i < b.N; i++ {
		if res := Add(a1, a2); res != c {
			b.Errorf("add(%d, %d) result: %d, expect: %d", a1, a2, res, c)
		}
	}
}

func TestNonRepeatSubString(t *testing.T) {
	cases := []struct {
		str    string
		length int
	}{
		{"a", 1},
		{"abcab", 3},
		{"灰化肥会挥发", 6},
		{"", 0},
	}
	for _, c := range cases {
		if res := NonRepeatSubStringMaxLen(c.str); res != c.length {
			t.Errorf("%s result: %d, expect: %d", c.str, res, c.length)
		}
	}
}

// go test -bench . -cpuprofile=cpu.out benchmark分析耗时
// go tool pprof cpu.out
func BenchmarkNonRepeatSubString(b *testing.B) {
	str := "灰化肥会挥发"
	for i := 0; i < 15; i++ {
		str = str + str
	}
	b.Logf("len(str) = %d", len(str))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res := NonRepeatSubStringMaxLen1(str)
		if res != 6 {
			b.Errorf("result: %d, expect: %d", res, 6)
		}
	}
}
