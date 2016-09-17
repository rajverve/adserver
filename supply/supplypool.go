package supply

import (
	"log"
)

type SupplyPool struct {
	ch chan *Supply
}

func NewSupplyPool(size int) *SupplyPool {
	return &SupplyPool {
		ch: make(chan *Supply, size),
	}
}

func (pool *SupplyPool) GetResource() *Supply {
	select {
	case r := <-pool.ch:
		return r
	default:
		log.Println("Creating new supply")
		return &Supply{
			Decision: make(chan bool),
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



