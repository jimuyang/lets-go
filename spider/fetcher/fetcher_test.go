package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	got, err := Fetch("http://www.zhenai.com/zhenghun/aba")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", got)
}
