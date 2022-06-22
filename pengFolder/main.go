package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting peng server")
	handleRequests()
}

func handleRequests() {
	fmt.Println("Peng is up and running")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/peng", pengHandler).Methods("GET")
	router.HandleFunc("/peng/ping", pengPingHandler).Methods("GET")
	router.HandleFunc("/peng/pong", pengPongHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8082", router))
}

func pengHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved request. Peng says hello")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := "Peng response written"
	bodyBytes := []byte(res)
	w.Write(bodyBytes)
	fmt.Println("Response written")
}

func pengPingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved request. Peng hits ping")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := "Peng Ping response written"
	bodyBytes := []byte(res)
	w.Write(bodyBytes)
	fmt.Println("Response written")
}

func pengPongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved request. Peng hits pong")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := "Peng Pong response written"
	bodyBytes := []byte(res)
	w.Write(bodyBytes)
	fmt.Println("Response written")
}
