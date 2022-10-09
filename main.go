package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {

	type PostMedia struct {
		OrderTicketList []int `json:"order_ticket_list,omitempty"`
	}
	s := PostMedia{}
	for _, q := range s.OrderTicketList {
		spew.Dump(q)
	}
}
