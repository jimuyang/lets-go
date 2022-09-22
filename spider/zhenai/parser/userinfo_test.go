package parser

import (
	"fmt"
	"github.com/jimuyang/lets-go/spider/fetcher"
	"testing"
)

func TestParseUserInfo(t *testing.T) {
	fetch, _ := fetcher.Fetch("https://album.zhenai.com/u/1402882293")
	fmt.Println(ParseUserInfo(fetch))
}
