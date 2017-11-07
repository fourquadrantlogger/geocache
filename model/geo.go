package common

import "math"

//Location 位置
type Location struct {
	//Lng 经度
	Lng float64
	//Lat 纬度
	Lat float64
}

//EarchR 地球半径
const EarchR = 6370856

//GEOdistance
// R·arc cos[cosβ1cosβ2cos（α1-α2）+sinβ1sinβ2]
func GEOdistance(p Location, ps Location) (distance float64) {

	distance = EarchR * math.Acos(math.Cos(ps.Lng)*math.Cos(ps.Lat)*2*math.Cos(p.Lat-p.Lng)+math.Sin(ps.Lng)*math.Sin(ps.Lat))
	return
}

//GEOdistances
// R·arc cos[cosβ1cosβ2cos（α1-α2）+sinβ1sinβ2]
func GEOdistances(p Location, ps map[int]Location) (distance map[int]float64) {
	for k, l := range ps {
		distance[k] = GEOdistance(p, l)
	}
	return
}
