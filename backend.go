package main

import (
  "time"
	"log"
	"net/http"
	"github.com/gorilla/mux"
  "github.com/ReneKroon/ttlcache/v2" 
)

func handleRequests() {
  r := mux.NewRouter().StrictSlash(true);
  r.HandleFunc("/busstop/", busTimingsAtBusStop).Methods("GET")
  r.HandleFunc("/busline/", busLocationInBusLine).Methods("GET")
  log.Fatal(http.ListenAndServe(":3000",r)) 
}

// Generating a cache
var cache ttlcache.SimpleCache = ttlcache.NewCache()

func main () {
// Data persists in the cache for 1s
  cache.SetTTL(time.Duration(1 * time.Second))
  handleRequests()
}
