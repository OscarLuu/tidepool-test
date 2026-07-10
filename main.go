package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.Handle("/", http.FileServer(http.Dir(".")))

	// Serve the tic-tac-toe game at a clean URL.
	http.HandleFunc("/tictactoe", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "tic-tac-toe.html")
	})

	addr := ":" + port
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
