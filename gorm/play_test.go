package gorm

import (
	"strings"
	"testing"
)

func TestUseGorm(t *testing.T) {
	ran := strings.TrimSpace("[9900, )")
	if len(ran) < 2 {
		t.Log("111")
	}
	first, last := ran[0], ran[len(ran)-1]
	mids := strings.Split(ran[1:len(ran)-1], ",")
	if len(mids) != 2 {
		t.Log("222")
	}
	lower, upper := strings.TrimSpace(mids[0]), strings.TrimSpace(mids[1])
	if lower != "" {
		t.Logf("q.From(%v).IncludeLower(%v)", lower, first == '[')
	}
	if upper != "" {
		t.Logf("q.To(%v).IncludeUpper(%v)", upper, last == ']')
	}
}
