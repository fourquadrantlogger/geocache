package common

import (
	"math"
)

//Location 位置
type Location struct {
	//Lat 经度
	Lat float64
	//Lng 纬度
	Lng float64
}

func (l Location) lng() float64 {
	return math.Pi * l.Lng / 180.0
}
func (l Location) lat() float64 {
	return math.Pi * l.Lat / 180.0
}

//EarchR 地球半径
const EarchR = 6371171

//GEOdistance 地球2点间距离
func GEOdistance(p Location, ps Location) (distance float64) {
	angle := GEOangle(p, ps)
	distance = EarchR * angle
	return
}

//GEOangle 地球2点间弧度
func GEOangle(p Location, ps Location) float64 {
	lat1 := p.lat()
	lng1 := p.lng()
	lat2 := ps.lat()
	lng2 := ps.lng()

	lat := lat2 - lat1
	lng := lng2 - lng1

	d1 := math.Sin(lat*0.5) * math.Sin(lat*0.5)
	d2 := math.Cos(lat1) * math.Cos(lat2) * math.Sin(lng*0.5) * math.Cos(lat1) * math.Cos(lat2) * math.Sin(lng*0.5)

	return 2 * math.Asin(math.Sqrt(d1+d2))
}

//GEOdistances 地球n点间距离p的距离
// R·arc cos[cosβ1cosβ2cos（α1-α2）+sinβ1sinβ2]
func GEOdistances(p Location, ps map[int]Location) (distance map[int]float64) {
	distance = make(map[int]float64)
	for k, l := range ps {
		distance[k] = GEOdistance(p, l)
	}
	return
}
