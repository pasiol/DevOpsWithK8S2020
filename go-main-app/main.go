package main

import (
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	uuid := uuid.New()
	localTime := time.Now()
	fmt.Printf("%v %s\n", localTime.UTC().Format(time.RFC3339Nano), uuid.String())
	fmt.Fprintf(w,"%s %s", localTime.UTC().Format(time.RFC3339Nano), uuid.String())

}

func main() {
	port := "0.0.0.0:3000"
	log.Printf("Web server main application starting at port %s.", port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(port, nil))
}


