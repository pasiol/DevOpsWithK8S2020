package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var counter = 0

func handler(w http.ResponseWriter, r *http.Request) {
	counter++
	fmt.Fprintf(w, "pong %d", counter)
	log.Printf("pong %d", counter)
}

func main() {
	address := "0.0.0.0:" + os.Getenv("APP_PORT")

	log.Printf("go-pingopong starting at port %s.", address)

	http.HandleFunc("/pingpong", handler)
	log.Fatal(http.ListenAndServe(address, nil))

}
