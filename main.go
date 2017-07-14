package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/czerwonk/asn_lookup/bird"
	"github.com/gorilla/mux"
)

const version = "0.2"

var (
	showVersion   = flag.Bool("version", false, "Show version information")
	listenAddress = flag.String("listen.address", ":10179", "Address to listen for web service requests")
	birdSocket    = flag.String("bird.socket", "", "Socket to communicate with bird routing daemon")
	bird6Socket   = flag.String("bird6.socket", "", "Socket to communicate with bird6 routing daemon")
)

func main() {
	flag.Parse()

	if *showVersion {
		printVersion()
		os.Exit(0)
	}

	startServer()
}

func printVersion() {
	fmt.Println("asn_lookup")
	fmt.Printf("Version: %s\n", version)
	fmt.Println("Author(s): Daniel Czerwonk")
	fmt.Println("Webservice providing announced prefixes of an specific ASN based on routing daemon information")
}

func startServer() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/asn/{asn}", handleAsnRequest)
	r.HandleFunc("/asn", lookupAsnHandler)

	log.Printf("Starting server to listen on %s\n", *listenAddress)

	if len(*birdSocket) > 0 {
		log.Printf("bird socket: %s\n", *birdSocket)
	}

	if len(*bird6Socket) > 0 {
		log.Printf("bird6 socket: %s\n", *bird6Socket)
	}

	log.Fatal(http.ListenAndServe(*listenAddress, r))
}

func handleAsnRequest(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	a := vars["asn"]

	b := bird.New(*birdSocket, *bird6Socket)

	as, err := b.GetAs(a)
	if err != nil {
		log.Println(err)
		return
	}

	xml, err := as.ToXml()
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(xml)
}

func lookupAsnHandler(w http.ResponseWriter, r *http.Request) {
	ip := r.URL.Query().Get("ip")
	x := net.ParseIP(ip)

	if x == nil {
		http.Error(w, "No IP specified", 400)
		return
	}

	b := bird.New(*birdSocket, *bird6Socket)

	as, err := b.GetAsByIP(x)
	if err != nil {
		log.Println(err)
		return
	}

	xml, err := as.ToXml()
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(xml)
}
