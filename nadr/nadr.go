package nadr

import (
	"net/http"
	"fmt"
	"log"
)

var pool = make(chan *NADR, 10)

type NADR struct {
	Decision chan bool
	w        http.ResponseWriter
	req      *http.Request
}

func (n *NADR) Initialize(w http.ResponseWriter, req *http.Request) {
	n.w = w
	n.req = req
}

func GetResource() *NADR {
	select {
	case n := <-pool:
		return n
	default:
		log.Println("Requests exceeded max requests")
		return &NADR {
			Decision: make(chan bool),
		}
	}
}

func ReturnResource(n *NADR) {
	select {
	case pool <- n:
	// return n to free list
	default:
	// free list full, allow n to be garbage collected
	}
}


func (n *NADR) Decide() {
	n.Decision <- true
}

func (n *NADR) Bid() {
	fmt.Fprint(n.w, "Yes, I'm bidding")
}


func (n *NADR) Normalize() {

}