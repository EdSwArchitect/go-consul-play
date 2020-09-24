# go-consul-play

## Overview
This excercise does a few things:
1. Simply inserts a key/value into Consul
1. Simple get the value from the key put into Consul
1. Demonstrates unencrypted use or TLS encrypted used
1. Uses another Go API to flatten a JSON file and place the flattened key/value pairs into Consul

## Build instructions
1. Do 'go get github.com/opencopilot/consulkvjson' to get the API that can be used to flatten the JSON
1. Do 'go build -o go-consul-play' to build everything

## Run instructions

Start Consul in development mode. In order to do this, the command is done in the same directory as the certificates created

>  consul agent -config-file=config-tls.json  -dev -data-dir ~/consul-data

This will start a Consul agent, unsecured on port 8500, TLS on port 8501


Usage:

Usage of ./go-consul-play:

>
>   -config string
>
>     	The JSON file to load
>
>   -dc string
>
>     	the data center (default "dc1")
>
>  -loadJson
>
>     	Load the JSON file
>
>   -tls
>     	If true, use TLS, otherwise, normal
>

To load a configuration file:

> ./go-consul-play -config ~/Development/CYBER/Common.Config/talonConfig.json -loadJson=true -tls=true

To do the simple insert of a k/v and a read of the k/v

> ./go-consul-play -config ~/Development/CYBER/Common.Config/talonConfig.json -loadJson=false


