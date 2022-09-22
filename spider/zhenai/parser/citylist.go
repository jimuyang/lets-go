package parser

import (
	"regexp"
	"strings"

	"github.com/jimuyang/lets-go/spider/engine"
)

const cityListExp = "\"linkContent\":\"([^\"]+)\",\"linkURL\":\"(http:\\\\u002F\\\\u002Fwww.zhenai.com\\\\u002Fzhenghun\\\\u002F[a-z0-9]+)\""

// 解析征婚首页的城市列表
func ParseCityList(bytes []byte) ([]interface{}, []engine.Job) {
	var allNames []interface{}
	var newJobs []engine.Job

	cityListReg := regexp.MustCompile(cityListExp)
	cityList := cityListReg.FindAllSubmatch(bytes, -1)
	for _, u := range cityList {
		name, cityURL := string(u[1]), string(u[2])
		// 替换地址内的unicode码点
		link := strings.ReplaceAll(cityURL, "\\u002F", "/")
		allNames = append(allNames, name)
		newJobs = append(newJobs, engine.Job{
			URL:    link,
			Parser: ParseCityUserList,
		})
	}
	return allNames, newJobs
}
