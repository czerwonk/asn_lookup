package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/asn/{asn}", HandleAsnRequest)
}

func HandleAsnRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	asn := vars["asn"]

	return nil
}
