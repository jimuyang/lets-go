package main

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"testing"
)

func TestGraph_Draw(t *testing.T) {
	//rand.Seed(time.Now().Unix())
	g := &Graph{
		Name:   "test",
		Points: []*Point{{Name: "start"}, {Name: "mid"}, {Name: "end"}},
		Connections: []*Connection{
			{
				Name: "start_to_mid",
				From: "start",
				To:   "mid",
			},
			{
				Name: "mid_to_end",
				From: "mid",
				To:   "end",
			},
		},
		drawPointFunc: map[string]DrawPoint{
			"start": demoGetData,
			"end":   demoGetData,
		},
		drawConnectionFunc: map[string]DrawConnection{
			"start_to_mid": demoRelation,
			"mid_to_end":   demoRelation,
		},
	}
	graphData := g.Draw([]*PointData{{
		PointName: "start",
		DataId:    "0",
	}})

	a, _ := json.Marshal(graphData)
	t.Log(string(a))
}

func demoGetData(from []*PointData) []*PointData {
	res := make([]*PointData, 0)
	for _, f := range from {
		res = append(res, &PointData{
			DataId: f.DataId,
			Data:   "demo data",
		})
	}
	return res
}

func demoRelation(from []*PointData) map[string][]*PointData {
	res := make(map[string][]*PointData)
	for _, f := range from {
		for i := 0; i < rand.Intn(5); i++ {
			res[f.DataId] = append(res[f.DataId], &PointData{
				DataId: strconv.Itoa(100000 + rand.Intn(100000)),
			})
		}
	}
	return res
}
