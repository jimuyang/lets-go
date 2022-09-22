package parser

import (
	"github.com/jimuyang/lets-go/spider/engine"
	"log"
	"regexp"
)

const memberURLPrefix = "https://album.zhenai.com/u/"

var memberListRe = regexp.MustCompile("\"memberList\":\\[([^]]+)")

var memberIDRe = regexp.MustCompile("\"memberID\":(.*?),")

// 解析城市信息下的用户列表
func ParseCityUserList(bytes []byte) ([]interface{}, []engine.Job) {
	memberBytes := matchAndFind(bytes, memberListRe)
	if memberBytes == nil {
		log.Println("cannot find members")
		return nil, nil
	}

	// member
	reg := regexp.MustCompile("\\{.*?}")
	members := reg.FindAll(memberBytes, -1)

	var results []interface{}
	var newJobs []engine.Job
	for _, member := range members {
		// get memberID
		memberID := matchAndFindString(member, memberIDRe)
		memberURL := memberURLPrefix + memberID
		newJobs = append(newJobs, engine.Job{
			URL:    memberURL,
			Parser: ParseUserInfo,
		})
		results = append(results, memberID)
	}
	return results, newJobs
}
