package parser

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/jimuyang/lets-go/spider/fetcher"
)

func TestParseCityUserList(t *testing.T) {
	fetch, _ := fetcher.Fetch("http://www.zhenai.com/zhenghun/aba")
	fmt.Println(ParseCityUserList(fetch))
}

var TemplateFillFormatRe = regexp.MustCompile("#{(.*?)}")

func Test_fillData(t *testing.T) {
	str := "你好 你的#{shopName}已开业#{ss}"
	data := make(map[string]string)
	data["shopName"] = "门店名称"
	data["ss"] = "sssss"
	matches := TemplateFillFormatRe.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		if value, ok := data[match[1]]; ok {
			str = strings.ReplaceAll(str, match[0], value)
		}
	}
	t.Log(str)
}
