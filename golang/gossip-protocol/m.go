package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/hashicorp/memberlist"
)

type Node struct {
	memberlist *memberlist.Memberlist
}

type Item struct {
	Ip     string `json:"ip"`
	Status string `json:"status"`
}

func (n *Node) handler(w http.ResponseWriter, req *http.Request) {

	var items []Item

	fmt.Println(n.memberlist.Members())
	for _, member := range n.memberlist.Members() {
		timeOut := time.Duration(time.Second*5) * time.Second
		_, err := net.DialTimeout("tcp", member.FullAddress().Addr, timeOut)

		fmt.Println(member)
		if err != nil {
			items = append(items, Item{Ip: member.FullAddress().Addr, Status: "DOWN"})
		} else {
			items = append(items, Item{Ip: member.FullAddress().Addr, Status: "UP"})
		}
	}

	js, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
