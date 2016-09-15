package supply

import (
	"log"
)

var pool = make(chan *Supply, 10)


func GetResource() *Supply {
	select {
	case s := <-pool:
		return s
	default:
		log.Println("Requests exceeded max requests")
		return &Supply{
			Decision: make(chan bool),
		}
	}
}

func ReturnResource(s *Supply) {
	select {
	case pool <- s:
	// return n to free list
	default:
	// free list full, allow n to be garbage collected
	}
}



