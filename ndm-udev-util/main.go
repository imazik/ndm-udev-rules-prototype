package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	url    = "http://localhost:8080/diskevent"
)

func main() {
	diskDetails := make(map[string]string)
	requestMethod := os.Args[1]
	inputArgs := os.Args[2:]
	for _, arg := range inputArgs {
		parts := strings.Split(arg, "=")
		if len(parts) == 2 {
			diskDetails[parts[0]] = parts[1]
		} else if len(parts) == 1 {
			diskDetails[parts[0]] = ""
		}
	}
	jsonString, err := json.Marshal(diskDetails)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	client := &http.Client{}
	req, err := http.NewRequest(requestMethod, url, bytes.NewBuffer(jsonString))
	if err != nil {
		log.Fatal("Cannot create request", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client.Do(req)
}
