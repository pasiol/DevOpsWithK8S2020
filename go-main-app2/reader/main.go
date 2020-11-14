package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
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

func readLines(filename string) string {
	f, err := os.Open(os.Getenv("APP_INPUT_FILE"))
	if err != nil {
		log.Fatalf("Opening file %s failed: %s", os.Getenv("APP_INPUT_FILE"), err)
	}
	scanner := bufio.NewScanner(f)
	content := ""
	for scanner.Scan() {
		line := scanner.Text()
		content = content + line + "\n"
	}
	log.Printf("Readed file %s", os.Getenv("APP_INPUT_FILE"))
	return content
}

func handler(w http.ResponseWriter, r *http.Request) {
	lines := ""
	if fileExists(os.Getenv("APP_INPUT_FILE")) {
		lines = readLines(os.Getenv("APP_INPUT_FILE"))
	} else {
		log.Printf("Cannot find input file: %s", os.Getenv("APP_INPUT_FILE"))
	}
	fmt.Fprintf(w, "%s", lines)
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
	log.Printf("Starting reader: %v", start)

	address := "0.0.0.0:" + os.Getenv("APP_PORT")
	log.Printf("Web server go-main-app2 starting at address %s.", address)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(address, nil))
}
