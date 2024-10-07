package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	proxyReq, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
	if err != nil {
		http.Error(w, "Error creating proxy request", http.StatusInternalServerError)
		return
	}

	for name, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(name, value)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Error sending proxy request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	w.WriteHeader(resp.StatusCode)

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Printf("Error copying response body: %v", err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./proxy <port>")
		os.Exit(1)
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil || port < 1 || port > 65535 {
		fmt.Println("Invalid port number. Please use a number between 1 and 65535.")
		os.Exit(1)
	}

	http.HandleFunc("/", handleRequest)
	
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting proxy server on %s\n", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
