package main

import "net/http"

func main () {
	server := http.NewServeMux()
	server.HandleFunc("/ping", handlePing)

	if err := http.ListenAndServe(":8000", server); err != nil {
		panic(err)
	}
}

func handlePing (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}