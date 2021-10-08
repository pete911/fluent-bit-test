package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var port = 8080

func main() {

	if portEnv, ok := os.LookupEnv("FBT_PORT"); ok {
		if portInt, err := strconv.Atoi(portEnv); err == nil {
			port = portInt
		}
	}
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		indexPostHandler(w, r)
		return
	}

	io.Copy(io.Discard, r.Body)
	r.Body.Close()
	w.WriteHeader(http.StatusAccepted)
}

func indexPostHandler(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("read request: %v", err)))
		return
	}

	defer r.Body.Close()
	log.Println(string(b))
	w.WriteHeader(http.StatusCreated)
}
