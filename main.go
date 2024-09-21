package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var port string
var peers string

func main() {

	flag.StringVar(&port, "port", ":8080", "HTTP server port")
	flag.StringVar(&peers, "peers", "", "Comma-separated list of peer addresses")
	flag.Parse()

	peerList := strings.Split(peers, ",")
	nodeID := fmt.Sprintf("%s%d", "node", rand.Intn(100))

	cs := NewCacheServer(5, 1*time.Minute, nodeID, append([]string{"self"}, peerList...))

	http.HandleFunc("/get", cs.GetHandler)
	http.HandleFunc("/set", cs.SetHandler)

	fmt.Println("Starting server on port ", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
	}
}
