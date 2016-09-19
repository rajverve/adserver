package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/rajverve/adserver/supply"
	"google.golang.org/grpc"
	pb "github.com/rajverve/protobuf"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

var pool = supply.NewSupplyPool(10)

func main() {
	http.Handle("/adserver", http.HandlerFunc(processRequest))
	err := http.ListenAndServe(":55555", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	conn, err := grpc.Dial("4444", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSegmentationClient(conn)

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



