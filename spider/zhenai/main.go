package main

import (
	"fmt"
	"regexp"

	"github.com/jimuyang/lets-go/spider/engine"
	"github.com/jimuyang/lets-go/spider/zhenai/parser"
)

var thirdIDReg = regexp.MustCompile("id=([^&]*)")

func findThirdIDFromURL(url string) string {
	r := thirdIDReg.FindStringSubmatch(url)
	fmt.Println(r)
	if r == nil {
		return ""
	}
	return r[1]
}

func main() {
	engine.Run(engine.Job{
		//URL:    "http://www.zhenai.com/zhenghun",
		//Parser: parser.ParseCityList,
		URL:    "http://www.zhenai.com/zhenghun/aba",
		Parser: parser.ParseCityUserList,
	})
	// fmt.Println(findThirdIDFromURL("https://mapproxy.org/amap/v3/place/detail?id=B1235"))
}
