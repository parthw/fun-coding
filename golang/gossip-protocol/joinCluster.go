package main

import (
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

func JoinCluster(httpPort, clusterKey, knownIP, listenPort string) {
	msgCh := make(chan []byte)

	d := new(MyDelegate)
	d.msgCh = msgCh

	config := memberlist.DefaultLocalConfig()
	config.Name = "join:" + listenPort
	config.BindPort, _ = strconv.Atoi(listenPort)
	config.AdvertisePort, _ = strconv.Atoi(listenPort)
	config.SecretKey, _ = base64.StdEncoding.DecodeString(clusterKey)
	config.Delegate = d

	ml, err := memberlist.Create(config)

	if err != nil {
		panic(err)
	}

	node := Node{
		memberlist: ml,
	}

	_, err = ml.Join([]string{knownIP})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}
	fmt.Println(node.memberlist.Members())

	http.HandleFunc("/", node.handler)

	go func() {
		http.ListenAndServe(":"+httpPort, nil)
	}()

	tick := time.NewTicker(3 * time.Second)

	incomingSigs := make(chan os.Signal, 1)
	signal.Notify(incomingSigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, os.Interrupt)

	for {
		select {

		case <-tick.C:
			m := new(MyMessage)
			m.Key = "ping"
			m.Value = "ack"

			// ping to all
			for _, node := range ml.Members() {
				if node.Name == config.Name {
					continue // skip self
				}
				fmt.Println(node.FullAddress())
				log.Printf("send to %s msg: key=%s value=%s", node.Name, m.Key, m.Value)
				if err := ml.SendReliable(node, m.Bytes()); err != nil {
					log.Println(err)
				}
			}

		case data := <-d.msgCh:
			log.Println("GOT RECEIVED MESSAGE")
			msg, ok := ParseMyMessage(data)
			if !ok {
				continue
			}

			log.Printf("received msg: key=%s value=%s", msg.Key, msg.Value)
			log.Printf("number of nodes %d", ml.NumMembers())

		case rSignal := <-incomingSigs:
			fmt.Printf("\nRecevied Signal %s\n", rSignal)
			if err := ml.Leave(time.Second * 5); err != nil {
				fmt.Println(err)
			}
			os.Exit(0)
		}
	}
}
