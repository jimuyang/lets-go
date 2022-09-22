package helper

import (
	"fmt"
	"regexp"
	"strings"
)

// const cityListExp = "http:\\\\u002F\\\\u002Fwww.zhenai.com\\\\u002Fzhenghun\\\\u002F[a-z0-9]+"

const cityListExp = "\"linkContent\":\"([^\"]+)\",\"linkURL\":\"(http:\\\\u002F\\\\u002Fwww.zhenai.com\\\\u002Fzhenghun\\\\u002F[a-z0-9]+)\""

const leftSlashUnicode = "\\\\u002F"

var cityListReg *regexp.Regexp

func init() {
	fmt.Println("regexp init")
	// url.QueryUnescape()
	fmt.Println(cityListExp)
	cityListReg = regexp.MustCompile(cityListExp)
	// fmt.Println(cityListReg.MatchString(str))
}

type ZhenghunCity struct {
	CityName string
	CityURL  string
}

func FindCityList(bytes []byte) []ZhenghunCity {
	result := make([]ZhenghunCity, 0)
	cityList := cityListReg.FindAllSubmatch(bytes, -1)
	for _, u := range cityList {
		cityName, cityURL := string(u[1]), string(u[2])
		// 替换地址内的unicode码点
		link := strings.ReplaceAll(cityURL, "\\u002F", "/")
		result = append(result, ZhenghunCity{cityName, link})
	}
	return result
}
