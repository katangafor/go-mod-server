package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

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
	fmt.Fprintln(w, "Hello World!")
}
func getPort() (string, error) {
	// the PORT is supplied by Heroku
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
