package main

import (
	"log"
	"net/http"
)

func wasmHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/wasm")
	http.ServeFile(w, r, "../../../wasmtest/main.wasm")
}
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("../../../wasmtest")))
	mux.HandleFunc("/main.wasm", wasmHandler)
	log.Fatal(http.ListenAndServe(":8001", mux))
}