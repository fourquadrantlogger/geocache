package common

import (
	"fmt"

	"testing"
)

func TestGEOdistance(t *testing.T) {
	beijing := Location{39.8999575057, 116.3871012193}  //北京
	shanghai := Location{31.2081296333, 121.4813465878} //上海

	fmt.Printf("GEOangle=%f\n", GEOangle(beijing, shanghai))
	fmt.Printf("GEOdistance=%f\n", GEOdistance(beijing, shanghai))

	//output
	//1035.722
}
