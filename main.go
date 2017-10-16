package main

import (
	"flag"
)

var (
	host      string
	port      int
	workerID  string
	requestID string
)

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "RPC Server Host")
	flag.IntVar(&port, "port", 1000, "RPC Server Port")
	flag.StringVar(&workerID, "workerId", "100a", "RPC Server Worker ID")
	flag.StringVar(&requestID, "requestId", "1", "RPC Server Request ID")
}

func main() {
	// Create connection
	functionrpc.
	// Open Channel
}
