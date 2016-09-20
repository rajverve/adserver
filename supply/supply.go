package supply

import (
	"encoding/json"
	"fmt"
	"github.com/rajverve/adserver/admatcher"
	"github.com/rajverve/protobuf"
	"golang.org/x/net/context"
	"io"
	"net/http"
)

type Supply struct {
	Decision        chan bool
	segClient       segmentation.SegmentationClient
	blacklistClient segmentation.BlacklistClient
	w               http.ResponseWriter
	req             *http.Request
}

func adRequestForSupply(p []byte) (a *segmentation.AdRequest, err error) {
	a = new(segmentation.AdRequest)
	err = json.Unmarshal(p, a)

	if err != nil {
		fmt.Println("Error unmarshaling json", err)
		return a, err
	}

	fmt.Printf("Received AdRequest %v\n", a)

	return a, err
}

func (s *Supply) segmentInfo(a *segmentation.AdRequest, ch chan *segmentation.SupplySegment) {
	seg, err := s.segClient.GetSegmentInfo(context.Background(), a)

	if err != nil {
		fmt.Println("Error getting segment info from vls")
		seg = nil
	}

	ch <- seg
}

func (s *Supply) blacklistInfo(a *segmentation.AdRequest, ch chan *segmentation.BlacklistResponse) {
	blacklist, err := s.blacklistClient.GetBlacklistInfo(context.Background(), a)

	if err != nil {
		fmt.Println("Error getting blacklist info from Blacklist Server")
		blacklist = nil
	}

	ch <- blacklist
}

func (s *Supply) Decide() {
	p := s.readRequest()
	a, err := adRequestForSupply(p)

	if err != nil {
		s.Decision <- false
		return
	}

	segChannel := make(chan *segmentation.SupplySegment)
	blacklistChannel := make(chan *segmentation.BlacklistResponse)

	go s.segmentInfo(a, segChannel)
	go s.blacklistInfo(a, blacklistChannel)

	seg, blacklist := <-segChannel, <-blacklistChannel

	if seg == nil || blacklist == nil || blacklist.Blacklist == true {
		s.Decision <- false
		return
	}

	ads := admatcher.AdsForSegment(seg)

	if len(ads) > 0 {
		s.Decision <- true
	} else {
		s.Decision <- false
	}
}

func (s *Supply) Bid() {
	fmt.Fprint(s.w, "Yes, I'm bidding")
}

func (s *Supply) readRequest() []byte {
	p := make([]byte, 256)
	read, err := s.req.Body.Read(p)

	if err != io.EOF && err != nil {
		fmt.Printf("supply: Error reading request %v\n", err)
	}

	return p[:read]
}
