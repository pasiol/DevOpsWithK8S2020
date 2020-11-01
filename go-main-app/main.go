package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main() {

	for true {
		uuid := uuid.New()
		localTime := time.Now()

		fmt.Printf("%v %s\n", localTime.UTC().Format(time.RFC3339Nano), uuid.String())
		time.Sleep(5 * time.Second)
	}
}
