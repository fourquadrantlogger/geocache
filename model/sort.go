package common

import "sort"

//LocationInfo 带信息的位置
type LocationInfo struct {
	Info map[string]interface{}
	Location
}

//GEOKNN 根据球面
func GEOKNN(p Location, ps map[int]Location) (ls []LocationInfo) {
	distances := GEOdistances(p, ps)
	for k, v := range ps {
		l := *(new(LocationInfo))
		l.Location = v
		l.Info = make(map[string]interface{})
		l.Info["index"] = k
		l.Info["distance"] = distances[k]
		ls = append(ls, l)
	}

	sort.Slice(ls, func(i, j int) bool { return ls[i].Info["distance"].(float64) < ls[j].Info["distance"].(float64) })
	return
}
