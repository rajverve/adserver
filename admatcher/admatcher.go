package admatcher

import (
	"fmt"
	"github.com/rajverve/protobuf"
	"reflect"
	"sync"
)

type Ad struct {
}

type AdMatcher struct {
	ads map[segmentation.SupplySegment]Ad
	mux sync.Mutex
}

var admatcher = AdMatcher{ads: make(map[segmentation.SupplySegment]Ad)}

func AdsForSegment(s *segmentation.SupplySegment) []Ad {
	admatcher.mux.Lock()
	found := findMatch(s, new(segmentation.SupplySegment), 0)
	admatcher.mux.Unlock()

	return found
}

func findMatch(test *segmentation.SupplySegment, current *segmentation.SupplySegment, index int) []Ad {
	fmt.Printf("To match %v\n", test)

	testVal := reflect.Indirect(reflect.ValueOf(test)).Field(index)

	fmt.Printf("Current Before %v\n", current)
	// set current's value equal to test's value
	reflect.Indirect(reflect.ValueOf(current)).Field(index).Set(testVal)

	// look up in map

	fmt.Printf("Current After %v\n", current)

	return make([]Ad, 0, 10)
}
