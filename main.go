package main

import (
	"encoding/json"
)

func main() {

	// rankResp := reqRank()
	// poiIds := ParseRank(rankResp)
	// ParseItem(reqItem(poiIds))
}

func ToJSON(v interface{}) string {
	ms, _ := json.Marshal(v)
	return string(ms)
}
