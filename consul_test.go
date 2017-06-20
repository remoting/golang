package main

import (
	"log"
	"os"
	"testing"
	"time"

	"fmt"

	"github.com/hashicorp/consul/api"
)

func TestConsul002(t *testing.T) {
	sigs := make(chan os.Signal, 1)
	NewConsulClient([]string{"192.168.31.240:8500"}, "http", "", "", "")
	go func() {
		var index uint64 = 0
		for {
			fmt.Printf("\n%d===========\n", index)
			index = watcher(index)
		}
	}()
	<-sigs
}

func watcher(index uint64) uint64 {
	keys, meta, err := Client.Catalog().Service("jira-1", "", &api.QueryOptions{
		WaitTime:  5 * time.Second,
		WaitIndex: index,
	})
	//fmt.Printf("%v", index)
	//fmt.Printf("watcher...")

	if err != nil {
		log.Fatal(err)
		return 0
	}

	for _, key := range keys {
		//if key.ModifyIndex > index {
		fmt.Printf("KeyIndex: %s Value: %s,%s,%v,%v \n", key.ModifyIndex, key.ServiceID, key.ServiceTags, key.ServiceAddress, key.ServicePort)
		//fmt.Printf("Services %v", key)

		//}
	}
	fmt.Printf("\nmeta,%v======", meta.LastIndex)
	return meta.LastIndex

}
func TestConsul(t *testing.T) {
	NewConsulClient([]string{"192.168.31.240:8500"}, "http", "", "", "")
	stopChan := make(chan bool, 1)
	for {
		resp, err := service(stopChan)
		if err != nil {
			fmt.Printf("%v", err.Error)
			time.Sleep(time.Second * 2)
			continue
		}
		fmt.Printf("%v", resp)
	}
}

type watchResponse struct {
	waitIndex uint64
	err       error
	data      map[string][]string
	meta      *api.QueryMeta
}

func service(stopChan chan bool) (watchResponse, error) {
	var waitIndex uint64 = 10
	respChan := make(chan watchResponse)
	go func() {
		opts := api.QueryOptions{
			WaitIndex: waitIndex,
		}
		_data, meta, err := Client.Catalog().Services(&opts)
		if err != nil {
			respChan <- watchResponse{waitIndex, err, _data, meta}
			return
		}
		respChan <- watchResponse{meta.LastIndex, err, _data, meta}
	}()
	select {
	case <-stopChan:
		return watchResponse{0, nil, nil, nil}, nil
	case r := <-respChan:
		return r, nil
	}
}
