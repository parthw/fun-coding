package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hashicorp/consul/api"
)

var isLeader bool

func GetConsulClient() *api.Client {

	config := api.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("client err: %v", err)
	}
	return client
}

func main() {
	// ttl in seconds
	ttl := 10
	ttlS := fmt.Sprintf("%ds", ttl)
	serviceKey := "service/distributed-lock/leader"
	serviceName := "distributed-lock"

	// build client
	client := GetConsulClient()
	// create session
	sEntry := &api.SessionEntry{
		Name:      serviceName,
		TTL:       ttlS,
		LockDelay: 1 * time.Millisecond,
	}
	sID, _, err := client.Session().Create(sEntry, nil)
	if err != nil {
		log.Fatalf("session create err: %v", err)
	}

	// auto renew session
	doneCh := make(chan struct{})
	go func() {
		err = client.Session().RenewPeriodic(ttlS, sID, nil, doneCh)
		if err != nil {
			log.Fatalf("session renew err: %v", err)
		}
	}()

	log.Printf("Consul session : %+v\n", sID)

	// Lock acquisition loop
	go func() {
		acquireKv := &api.KVPair{
			Session: sID,
			Key:     serviceKey,
			Value:   []byte("node2"),
		}

		for {
			if !isLeader {
				acquired, _, err := client.KV().Acquire(acquireKv, nil)
				if err != nil {
					log.Fatalf("kv acquire err: %v", err)
				}

				if acquired {
					isLeader = true
					log.Printf("I'm the leader !\n")
				}
			}

			time.Sleep(time.Duration(ttl/2) * time.Second)
		}
	}()

	// wait for SIGINT or SIGTERM, clean up and exit
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
	close(doneCh)
	log.Printf("Destroying session and leaving ...")
	_, err = client.Session().Destroy(sID, nil)
	if err != nil {
		log.Fatalf("session destroy err: %v", err)
	}
	os.Exit(0)
}
