package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// usage:
// $ curl -X GET http://127.0.0.1:8282/ping
// OK

func main() {
	// create and set up gorilla/mux router
	r := mux.NewRouter()

	r.HandleFunc("/ping", handlePing)

	// set listen address
	const addr = "127.0.0.1:8282"

	// log that we're about to start the web server
	log.Println(fmt.Sprintf("listening on %s", addr))

	// listen on the address and serve requests according to the router 'r'.
	// press Ctrl+C to quit
	log.Fatal(http.ListenAndServe(addr, r))
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK"))
}
