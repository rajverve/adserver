package main

import (
	"fmt"
	"github.com/rajverve/adserver/supply"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"github.com/rajverve/adserver/admatcher"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

var pool *supply.SupplyPool
var listenPort = ":55555"
var vlsPort = ":4444"
var blacklistPort = ":3333"

func main() {

	admatcher.SeedAdMatcher() // for testing only

	vlsConn, err := grpc.Dial(vlsPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to vls server: %v", err)
		return
	}
	defer vlsConn.Close()

	blacklistConn, err := grpc.Dial(blacklistPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to blacklist server: %v", err)
		return
	}
	defer blacklistConn.Close()

	pool = supply.NewSupplyPool(10, vlsConn, blacklistConn)

	http.Handle("/adserver", http.HandlerFunc(processRequest))
	err = http.ListenAndServe(listenPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
		return
	}
}

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	f(w, req)
}

func processRequest(w http.ResponseWriter, req *http.Request) {
	d := pool.GetResource(w, req)

	go d.Decide()

	if shouldBid := <-d.Decision; shouldBid {
		d.Bid()
		fmt.Println("Put in a bid")
	} else {
		fmt.Println("Going to sit this one out")
	}

	pool.ReturnResource(d)
}
