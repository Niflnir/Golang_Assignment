package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	responses "uwave/structs"
  "github.com/ReneKroon/ttlcache/v2" 
)

func busLocationInBusLine(w http.ResponseWriter, r *http.Request) {
  allBusLines := []string{"44478", "44479", "44480", "44481"}
  // check for param key /?buslineid=busLineID
  busLineID, ok:= r.URL.Query()["buslineid"]
  if !ok || len(busLineID[0]) < 1 || !contains(allBusLines, busLineID[0]){
    log.Println("Url Param 'key' is missing")
    w.WriteHeader(400)
    return
  }
  buslineid:= busLineID[0]
  
  // check cache for corresponding buslineid
  val, err := cache.Get(buslineid)
  if err != ttlcache.ErrNotFound{
    json.NewEncoder(w).Encode(val) 
    return
  }

  url := "http://dummy.uwave.sg/busline/" + buslineid
  res, err := http.Get(url)
  if(err != nil){
     log.Fatal(err) 
  }

  // data read is in byte slice
  data, err := ioutil.ReadAll(res.Body)
  if err != nil {
     log.Fatal(err)
  }

  // defined struct to store the json response 
  var busLine responses.BusLine
  json.Unmarshal(data, &busLine) 

  var busLocationResponses []responses.BusLocationResponse 
  for i := range busLine.Vehicles {
    var busLocationResponse responses.BusLocationResponse
    busLocationResponse.BusLineName = busLine.BusLineName
    busLocationResponse.Latitude = busLine.Vehicles[i].Latitude 
    busLocationResponse.Longtitude = busLine.Vehicles[i].Longtitude
    busLocationResponse.Vehicle_id = busLine.Vehicles[i].Vehicle_id
    busLocationResponses = append(busLocationResponses, busLocationResponse)
  }
  cache.Set(buslineid, busLocationResponses)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)    
  json.NewEncoder(w).Encode(busLocationResponses) // Writes to the localhost page 
}

func contains(elems []string, v string) bool {
    for _, s := range elems {
        if v == s {
            return true
        }
    }
    return false
}

