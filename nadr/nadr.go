package nadr

import (
	"fmt"
	"encoding/json"
)

type NADR struct {
	deviceId string
	lat float64
	lng float64
}


func NewNadr(b []byte) *NADR {
	n := NADR{}
	err := json.Unmarshal(b, &n)

	if err {
		fmt.Println("Error unmarshaling json")
		return &NADR{}
	}

	return n;
}
