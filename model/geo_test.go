package common

import (
	"fmt"
	"math"
	"testing"
)

func TestGEOdistance(t *testing.T) {
	beijing := Location{116.5429835110, 39.7072013727}  //北京,39.7072013727,
	shanghai := Location{121.5307607342, 31.1658172507} //上海,

	fmt.Printf("%f\n", (180/math.Pi)*GEOangle(beijing, shanghai))
	fmt.Printf("%f\n", GEOdistance(beijing, shanghai))
	//
	//fmt.Printf("%f\n",math.Sin(math.Pi))
	//fmt.Printf("%f\n",math.Sin(math.Pi/4))
	//fmt.Printf("%f\n",math.Sqrt(2)/2)

}
