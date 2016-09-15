package supply


import (
	"net/http"
	"fmt"
)

type Supply struct {
	Decision chan bool
	w        http.ResponseWriter
	req      *http.Request
}

func (s *Supply) Initialize(w http.ResponseWriter, req *http.Request) {
	s.w = w
	s.req = req
}

func (s *Supply) Decide() {
	s.Decision <- true
}

func (s *Supply) Bid() {
	fmt.Fprint(s.w, "Yes, I'm bidding")
}


