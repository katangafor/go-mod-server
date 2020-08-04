package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	// get the port
	port, err := getPort()
	if err != nil {
		log.Fatal(err)
	}
	// GET /
	http.HandleFunc("/", hello)
	// start the server
	log.Printf("Listening on %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	m := Message{"Max", "is a bro", 1294706395881547000}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, string(b))
}

func getPort() (string, error) {
	// the PORT is supplied by Heroku
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}
