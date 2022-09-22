package main

import "strings"

type Graph struct {
	Name        string
	Points      []*Point
	Connections []*Connection

	// transient
	relation           map[string][]string // from -> []connection->to
	distinctName       map[string]bool
	drawPointFunc      map[string]DrawPoint
	drawConnectionFunc map[string]DrawConnection
}

type DrawPoint func([]*PointData) []*PointData

type DrawConnection func([]*PointData) map[string][]*PointData

func (g *Graph) Draw(start []*PointData) *GraphData {
	res := &GraphData{
		Meta:              g,
		PointDataMap:      make(map[string][]*PointData),
		PointDataRelation: make(map[string]map[string][]string),
	}
	g.analyze()
	g.draw(start, res, make(map[string]bool))
	return res
}

func (g *Graph) draw(fromData []*PointData, res *GraphData, drew map[string]bool) {
	if len(fromData) == 0 {
		return
	}
	fromPoint := fromData[0].PointName
	if drawPoint := g.drawPointFunc[fromPoint]; drawPoint != nil {
		fromData = drawPoint(fromData)
		for _, fromDataItem := range fromData {
			fromDataItem.PointName = fromPoint
		}
	}
	res.PointDataMap[fromPoint] = append(res.PointDataMap[fromPoint], fromData...)
	for _, connectionTo := range g.relation[fromPoint] {
		temp := strings.Split(connectionTo, "->")
		connection, toPoint := temp[0], temp[1]
		if drawConnection := g.drawConnectionFunc[connection]; drawConnection != nil {
			fromToMap := make(map[string][]string)
			for fromDataId, toData := range drawConnection(fromData) {
				fromKey := fromPoint + ":" + fromDataId
				for _, toDataItem := range toData {
					toDataItem.PointName = toPoint
					fromToMap[fromKey] = append(fromToMap[fromKey], toDataItem.Key())
				}
				g.draw(toData, res, drew)
			}
			res.PointDataRelation[connection] = fromToMap
		}
		drew[connection] = true
	}
}

func (g *Graph) analyze() {
	g.relation = make(map[string][]string)
	for _, conn := range g.Connections {
		g.relation[conn.From] = append(g.relation[conn.From], conn.Name+"->"+conn.To)
	}
}

type Point struct {
	Name string
}

type Connection struct {
	Name     string
	From, To string
}

// 可传输/持久化的图数据
type GraphData struct {
	Meta              *Graph                         // 图的元数据
	PointDataMap      map[string][]*PointData        // pointName -> []pointData
	PointDataRelation map[string]map[string][]string // connectionName -> fromDataKey -> []toDataKey
}

type PointData struct {
	PointName string
	DataId    string      // 唯一id
	RawData   string      // json数据
	Data      interface{} // transient
}

func (p *PointData) Key() string {
	return p.PointName + ":" + p.DataId
}
