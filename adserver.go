package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/rajverve/adserver/supply"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

var pool = supply.NewSupplyPool(10)

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
	d := pool.GetResource()
	d.Initialize(w, req)

	go d.Decide()

	if shouldBid := <-d.Decision; shouldBid {
		d.Bid()
		fmt.Println("Put in a bid")
	} else {
		fmt.Println("Going to sit this one out")
	}

	pool.ReturnResource(d)
}



