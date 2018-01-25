package main

import (
	"fmt"
	"log"
	"net/http"
)

// HelloServer :hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello, world!")
}

func main() {
	http.HandleFunc("/", HelloServer)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
