package supply


import (
	"net/http"
	"fmt"
    "io"
    "github.com/rajverve/adserver/nadr"
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
    	p := s.readRequest()
    	n := nadr.NewNadr(p)
	fmt.Println(n)
	s.Decision <- true
}

func (s *Supply) Bid() {
	fmt.Fprint(s.w, "Yes, I'm bidding")
}

func (s *Supply) readRequest() []byte {
    	p := make([]byte, 256)
    	read, err := s.req.Body.Read(p)
    
    	if err != io.EOF  && err != nil {
        	fmt.Printf("supply: Error reading request %v\n", err)
    	}

    	return p[:read]
}

