package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting ping server")
	handleRequests()
}

func handleRequests() {
	fmt.Println("Ping is up and running")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ping", pingHandler).Methods("GET")
	router.HandleFunc("/ping/pong", pingPongHandler).Methods("GET")
	router.HandleFunc("/ping/peng", pingPengHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved request. Ping says hello")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := "Ping response written"
	bodyBytes := []byte(res)
	w.Write(bodyBytes)
	fmt.Println("Response written")
}

func pingPongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved request. Ping hits pong")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := "Ping Pong response written"
	bodyBytes := []byte(res)
	w.Write(bodyBytes)
	fmt.Println("Response written")
}

func pingPengHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved request. Ping hits peng")

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res := "Ping Peng response written"
	bodyBytes := []byte(res)
	w.Write(bodyBytes)
	fmt.Println("Response written")
}
