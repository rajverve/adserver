package nadr

import (
	"fmt"
	"encoding/json"
)

type NADR struct {
	DeviceId string
	Lat float64
	Lng float64
}


func NewNadr(b []byte) *NADR {
	fmt.Println(string(b))
	n := NADR{}
	err := json.Unmarshal(b, &n)

	if err != nil {
		fmt.Println("Error unmarshaling json", err)
		return &NADR{}
	}

	return &n;
}
