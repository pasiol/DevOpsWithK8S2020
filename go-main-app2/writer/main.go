package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func createDir(path string) {
	os.MkdirAll(path, os.ModePerm)
}

func createFile(filename string, content string) {
	if !fileExists(filepath.Dir(filename)) {
		createDir(filepath.Dir(filename))
	}
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatalf("Creating file %s failed: %s", filename, err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func appendString2File(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Cannot open output file %s %s", filename, err)
	}
	if _, err := f.Write([]byte(text)); err != nil {
		log.Fatalf("Cannot append string to file %s %s", filename, err)
	}
	if err := f.Close(); err != nil {
		log.Fatalf("Cannot close file %s %s", filename, err)
	}

}

func writeUUID() {
	if !fileExists(os.Getenv("APP_OUTPUT_FILE")) {
		createFile(os.Getenv("APP_OUTPUT_FILE"), "")
	}
	for true {
		uuid := uuid.New()
		localTime := time.Now()

		u := fmt.Sprintf("%v %s\n", localTime.UTC().Format(time.RFC3339Nano), uuid.String())
		appendString2File(os.Getenv("APP_OUTPUT_FILE"), u)
		time.Sleep(5 * time.Second)
	}
}

func main() {
	start := time.Now()
	if !fileExists(os.Getenv("APP_LOG_FILE")) {
		createFile(os.Getenv("APP_LOG_FILE"), "")
	}
	f, err := os.OpenFile(os.Getenv("APP_LOG_FILE"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Opening log file failed: %s", err)
	}
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
	log.Printf("Starting writer: %v", start)
	writeUUID()
}
