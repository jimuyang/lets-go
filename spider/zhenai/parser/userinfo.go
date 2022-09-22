package parser

import (
	"fmt"
	"github.com/jimuyang/lets-go/spider/engine"
	"regexp"
)

type UserInfo struct {
	id          string
	gender      string
	nickName    string
	description string
	monologue   string
	basicInfo   string
	detailInfo  string
	objectInfo  string
}

func (u *UserInfo) String() string {
	return fmt.Sprintf("{id: %s, name:%s, gender: %s, des: %s}", u.id, u.nickName, u.gender, u.description)
}

var idRe = regexp.MustCompile("ID：(.*?)<")
var genderRe = regexp.MustCompile("\"genderString\":\"(.*?)\"")
var nickNameRe = regexp.MustCompile("<span class=\"nickName\".*?>(.*?)</span>")
var descriptionRe = regexp.MustCompile("<div class=\"des f-cl\".*?>(.*?)</div>")
var basicInfoRe = regexp.MustCompile("\"basicInfo\":\\[([^]]+)")
var detailInfoRe = regexp.MustCompile("\"detailInfo\":\\[([^]]+)")
var objectInfoRe = regexp.MustCompile("\"objectInfo\":\\[([^]]+)")

// 解析用户信息
func ParseUserInfo(bytes []byte) ([]interface{}, []engine.Job) {
	u := &UserInfo{}
	u.id = matchAndFindString(bytes, idRe)
	u.gender = matchAndFindString(bytes, genderRe)
	u.nickName = matchAndFindString(bytes, nickNameRe)
	u.description = matchAndFindString(bytes, descriptionRe)
	u.basicInfo = matchAndFindString(bytes, basicInfoRe)
	u.detailInfo = matchAndFindString(bytes, detailInfoRe)
	u.objectInfo = matchAndFindString(bytes, objectInfoRe)
	return append([]interface{}{}, u), nil
}

func matchAndFindString(bytes []byte, exp *regexp.Regexp) string {
	find := matchAndFind(bytes, exp)
	if find == nil {
		return ""
	}
	return string(find)
}

func matchAndFind(bytes []byte, exp *regexp.Regexp) []byte {
	find := exp.FindSubmatch(bytes)
	if find == nil {
		return nil
	}
	return find[1]
}
