package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/rajverve/adserver/nadr"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

func main() {
	http.Handle("/adserver", http.HandlerFunc(processRequest))
	err := http.ListenAndServe(":55555", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	f(w, req)
}

func processRequest(w http.ResponseWriter, req *http.Request) {
	n := nadr.GetResource()
	n.Initialize(w, req)

	go n.Decide()

	if shouldBid := <-n.Decision; shouldBid {
		n.Bid()
		fmt.Println("Put in a bid")
	} else {
		fmt.Println("Going to sit this one out")
	}

	nadr.ReturnResource(n)
}



