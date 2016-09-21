package admatcher

import (
	"fmt"
	"github.com/rajverve/protobuf"
	"reflect"
	"sync"
)

type Ad struct {
	Name string
}

type AdMatcher struct {
	ads map[segmentation.SupplySegment]Ad
	mux sync.RWMutex
}

var admatcher = AdMatcher{ads: make(map[segmentation.SupplySegment]Ad)}

func AdsForSegment(s *segmentation.SupplySegment) []Ad {
	matched := make([]Ad, 0, 10)
	admatcher.mux.RLock()
	matched = findMatch(s, new(segmentation.SupplySegment), 0, matched)
	admatcher.mux.RUnlock()

	fmt.Println(matched)

	return matched
}

func findMatch(actual *segmentation.SupplySegment, current *segmentation.SupplySegment, index int, matched []Ad) []Ad {
	if index == reflect.Indirect(reflect.ValueOf(actual)).NumField() {
		return matched
	}

	next := *current
	actualFieldVal := reflect.Indirect(reflect.ValueOf(actual)).Field(index)
	reflect.ValueOf(&next).Elem().Field(index).Set(actualFieldVal)

	fmt.Printf("Index: %v, Testing: %v\n", index, next)
	if ad, ok := admatcher.ads[next]; ok {
		matched = append(matched, ad)
	}

	matched = findMatch(actual, current, index + 1, matched)
	matched = findMatch(actual, &next, index + 1, matched)

	return matched
}

func AddAd(segment *segmentation.SupplySegment, ad *Ad) {
	admatcher.mux.Lock()
	admatcher.ads[*segment] = *ad
	admatcher.mux.Unlock()
	fmt.Println(admatcher.ads)
}

func SeedAdMatcher() {
	ad1 := Ad{"ad1"}
	ad2 := Ad{"ad2"}

	seg1 := segmentation.SupplySegment{Audience:"Moms", PlaceName:"Target", PlaceType:"Mall"}
	seg2 := segmentation.SupplySegment{Audience:"Dads"}
	seg3 := segmentation.SupplySegment{PlaceType:"Pizza"}
	seg4 := segmentation.SupplySegment{Audience:"Kids", PlaceType:"Mall"}

	admatcher.ads[seg1] = ad1
	admatcher.ads[seg2] = ad1
	admatcher.ads[seg3] = ad2
	admatcher.ads[seg4] = ad2
}
