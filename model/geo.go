package common

import (
	"math"
)

//Location 位置
type Location struct {
	//Lng 经度
	Lng float64
	//Lat 纬度
	Lat float64
}

func (l Location) lng() float64 {
	return math.Pi * l.Lng / 180.0
}
func (l Location) lat() float64 {
	return math.Pi * l.Lat / 180.0
}

//EarchR 地球半径
const EarchR = 6378137

//GEOdistance 地球2点间距离
// R·arc cos[cosβ1cosβ2cos（α1-α2）+sinβ1sinβ2]
func GEOdistance(p Location, ps Location) (distance float64) {
	angle := GEOangle(p, ps)
	distance = EarchR * angle
	return
}

//GEOangle 地球2点间弧度
// sin(x1)*sin(x2)+cos(x1)*cos(x2)*cos(y1-y2)
func GEOangle(p Location, ps Location) float64 {
	x1 := p.lng()
	x2 := ps.lng()
	y1 := p.lat()
	y2 := ps.lat()
	//fmt.Println(x1,x2,y1,y2)
	//fmt.Println(math.Sin(x1)*math.Sin(x2))
	//fmt.Println(math.Cos(x1)*math.Cos(x2)*math.Cos(y1-y2))
	return math.Acos(math.Sin(x1)*math.Sin(x2) + math.Cos(x1)*math.Cos(x2)*math.Cos(y1-y2))
}

//GEOdistances 地球n点间距离p的距离
// R·arc cos[cosβ1cosβ2cos（α1-α2）+sinβ1sinβ2]
func GEOdistances(p Location, ps map[int]Location) (distance map[int]float64) {
	for k, l := range ps {
		distance[k] = GEOdistance(p, l)
	}
	return
}
