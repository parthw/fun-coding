package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/hashicorp/memberlist"
)

func InitGossipCluster(httpPort, listenPort int) {
	msgCh := make(chan []byte)
	d := new(MyDelegate)
	d.msgCh = msgCh

	config := memberlist.DefaultLocalConfig()

	clusterKey := make([]byte, 32)
	_, _ = rand.Read(clusterKey)
	fmt.Printf("new cluster key: %s\n", base64.StdEncoding.EncodeToString(clusterKey))
	config.Name = "init:" + strconv.Itoa(listenPort)
	config.BindPort = listenPort
	config.AdvertisePort = listenPort
	config.SecretKey = clusterKey
	config.Delegate = d

	ml, err := memberlist.Create(config)
	if err != nil {
		panic(err)
	}

	node := Node{memberlist: ml}

	http.HandleFunc("/", node.handler)
	go func() {
		http.ListenAndServe(":"+strconv.Itoa(httpPort), nil)
	}()

	// tick := time.NewTicker(time.Second * 3)
	incomingSigs := make(chan os.Signal, 1)
	signal.Notify(incomingSigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt)
	for {
		select {

		// case <-tick.C:
		// 	m := new(MyMessage)
		// 	m.Key = "ping"
		// 	m.Value = "ack"

		// 	// ping to all
		// 	for _, node := range ml.Members() {
		// 		if node.Name == config.Name {
		// 			continue // skip self
		// 		}
		// 		log.Printf("send to %s msg: key=%s value=%s", node.Name, m.Key, m.Value)
		// 		if err := ml.SendReliable(node, m.Bytes()); err != nil {
		// 			log.Println(err)
		// 		}
		// 	}

		case data := <-d.msgCh:
			msg, ok := ParseMyMessage(data)
			if !ok {
				continue
			}

			log.Printf("received msg: key=%s value=%s", msg.Key, msg.Value)
			log.Printf("number of nodes %d", ml.NumMembers())
			for _, node := range ml.Members() {
				if node.Name == config.Name {
					continue
				}
				fmt.Println(node.FullAddress())
				log.Printf("sending message %s %s", "pong", "ack")
				msg.Key = "pong"
				ml.SendReliable(node, msg.Bytes())
			}

		case rSignal := <-incomingSigs:
			fmt.Printf("\nRecevied Signal %s\n", rSignal)
			if err := ml.Leave(time.Second * 5); err != nil {
				fmt.Println(err)
			}
			os.Exit(0)
		}
	}
}
