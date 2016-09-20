package supply

import (
	"github.com/rajverve/protobuf"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type SupplyPool struct {
	ch            chan *Supply
	vlsConn       *grpc.ClientConn
	blacklistConn *grpc.ClientConn
}

func NewSupplyPool(size int, vlsConn *grpc.ClientConn, blacklistConn *grpc.ClientConn) *SupplyPool {
	return &SupplyPool{
		ch:            make(chan *Supply, size),
		vlsConn:       vlsConn,
		blacklistConn: blacklistConn,
	}
}

func (pool *SupplyPool) GetResource(w http.ResponseWriter, req *http.Request) *Supply {
	select {
	case r := <-pool.ch:
		r.w = w
		r.req = req
		return r
	default:
		log.Println("Creating new supply")
		return &Supply{
			Decision:        make(chan bool),
			segClient:       segmentation.NewSegmentationClient(pool.vlsConn),
			blacklistClient: segmentation.NewBlacklistClient(pool.blacklistConn),
			w:               w,
			req:             req,
		}
	}
}

func (pool *SupplyPool) ReturnResource(s *Supply) {
	select {
	case pool.ch <- s:
		// return n to free list
		log.Println("Returning Supply to pool")
	default:
		// free list full, allow n to be garbage collected
		log.Println("Allowing excess Supply to be garbage collected")
	}
}
