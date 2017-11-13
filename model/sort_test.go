package common

import (
	"fmt"
	"testing"
)

func TestGEOKNN(t *testing.T) {

	ps := make(map[int]Location)
	ps[0] = Location{39.8999575057, 116.3871012193} //北京
	ps[1] = Location{31.2081296333, 121.4813465878} //上海

	m := Location{34.2081296333, 118.4813465878} //我的位置
	fmt.Println(GEOKNN(m, ps))
}
