package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting pong server")
	handleRequests()
}

func handleRequests() {
	fmt.Println("Pong is up and running")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/pong", pongHandler).Methods("GET")
	router.HandleFunc("/pong/ping", pongPingHandler).Methods("GET")
	router.HandleFunc("/pong/peng", pongPengHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func pongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved request. Pong says hello")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := "Pong response written"
	bodyBytes := []byte(res)
	w.Write(bodyBytes)
	fmt.Println("Response written")
}

func pongPingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved request. Pong hits ping")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := "Pong Ping response written"
	bodyBytes := []byte(res)
	w.Write(bodyBytes)
	fmt.Println("Response written")
}

func pongPengHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved request. Pong hits peng")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := "Pong Peng response written"
	bodyBytes := []byte(res)
	w.Write(bodyBytes)
	fmt.Println("Response written")
}
