package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/consul/api"
	"github.com/opencopilot/consulkvjson"
)

var tls *bool
var loadJSON *bool
var dc *string
var configFile *string

func init() {
	log.Println("Initializing...")

	tls = flag.Bool("tls", false, "If true, use TLS, otherwise, normal")
	dc = flag.String("dc", "dc1", "the data center")
	loadJSON = flag.Bool("loadJson", false, "Load the JSON file")
	configFile = flag.String("config", "", "The JSON file to load")

	flag.Parse()

	log.Printf("Using TLS: %t\n", *tls)
	log.Printf("Data center: %s\n", *dc)
	log.Printf("Load JSON file: %t\n", *loadJSON)
	log.Printf("The config file: %s\n", *configFile)
}

func main() {
	config := api.DefaultConfig()

	if *tls {

		config.TLSConfig.CAFile = "/Uses/ebrown/certs2/consul-agent-ca.pem"
		config.TLSConfig.CertFile = "/Users/ebrown/certs2/dc1-client-consul-0.pem"
		config.TLSConfig.KeyFile = "/Users/ebrown/certs2/dc1-client-consul-0-key.pem"
		config.TLSConfig.InsecureSkipVerify = false
		config.TLSConfig.Address = "localhost:8501"
		config.Address = "localhost:8501"
		config.Scheme = "https"
	}

	config.Datacenter = *dc

	log.Printf("Config: %+v\n", config)

	// Get a new client
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	if !*loadJSON {
		// Get a handle to the KV API
		kv := client.KV()

		// PUT a new KV pair
		p := &api.KVPair{Key: "REDIS_MAXCLIENTS", Value: []byte("1000")}
		_, err = kv.Put(p, nil)
		if err != nil {
			panic(err)
		}

		// Lookup the pair
		pair, _, err := kv.Get("REDIS_MAXCLIENTS", nil)
		if err != nil {
			panic(err)
		}
		fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
	} else {
		loadJSONFile(client)
	}

}

// loadJSONFile load the json file, after flattened, into Consul
func loadJSONFile(client *api.Client) error {
	data, err := ioutil.ReadFile(*configFile)

	if err != nil {
		log.Printf("Error reading the file '%s'\n%s\n", *configFile, err)
		return err
	}

	kvs, err := consulkvjson.ToKVs(data)
	if err != nil {
		log.Printf("There was an error parsing the file: %v", err)
		return err
	}

	kv := client.KV()

	for _, keyValue := range kvs {

		p := &api.KVPair{Key: keyValue.Key, Value: []byte(keyValue.Value)}
		_, err = kv.Put(p, nil)

		if err != nil {
			log.Printf("There was an error putting the value '%s' : '%s', into Consul.\n%s\n", keyValue.Key, keyValue.Value, err)
			return err
		}
	}

	return nil

}
