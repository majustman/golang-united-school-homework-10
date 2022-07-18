package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", handleSuccess).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handleData)
	router.HandleFunc("/headers", handleHeaders)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	// host := os.Getenv("HOST")
	host := os.Getenv("localhost")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

func handleSuccess(w http.ResponseWriter, r *http.Request) {
	slicedURL := strings.Split(r.URL.String(), "/")
	msg := slicedURL[len(slicedURL)-1]
	fmt.Fprintf(w, "Hello, %s!", msg)
	return
}

func handleBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	return
}

func handleData(w http.ResponseWriter, r *http.Request) {
	d, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "I got message:\n%s", string(d))
}

func handleHeaders(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	b := r.Header.Get("b")
	nA, _ := strconv.Atoi(a)
	nB, _ := strconv.Atoi(b)
	res := strconv.Itoa(nA + nB)
	w.Header().Add("a+b", res)
}
