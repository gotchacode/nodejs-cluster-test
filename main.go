package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	t := time.Now()
	fmt.Println(t.String())
	fmt.Println("hello, world!")
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
