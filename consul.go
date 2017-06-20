package main

import "github.com/hashicorp/consul/api"

var conf *api.Config
var Client *api.Client

func NewConsulClient(nodes []string, scheme, cert, key, caCert string) error {
	conf = api.DefaultConfig()

	conf.Scheme = scheme

	if len(nodes) > 0 {
		conf.Address = nodes[0]
	}
	_client, err := NewClient()
	Client = _client
	if err != nil {
		return nil
	}
	return nil
}
func NewClient() (*api.Client, error) {
	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}
	/*
		kv := Client.KV()
		keys, _, _ := kv.List("mykey", nil)
		for _, key := range keys {
			fmt.Printf("Key: %s Value: %s", key.Key, key.Value)
		}
		fmt.Printf("%v", kv)
	*/
	return client, nil
}


