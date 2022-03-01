package main
import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	responses "uwave/structs"
  "github.com/ReneKroon/ttlcache/v2" 
)

func busTimingsAtBusStop(w http.ResponseWriter, r *http.Request) { 
  allBusStopIDs := []string{
  "378204", "383050", "378202", "383049", "382998", "378237", "378233", "378230",
  "378229", "378228", "378227", "382995", "378224", "378226", "383010", "383009",
  "383006", "383004", "378234", "383003", "378222", "383048", "378203", "382999", 
  "378225", "383014", "383013", "383011", "377906", "383018", "383015", "378207", 
  }
  
  // check for param key /?busid=busstopid
  busStopID, ok:= r.URL.Query()["busstopid"]
  if !ok || len(busStopID[0]) < 1|| !contains(allBusStopIDs, busStopID[0]) {
    log.Println("Url Param 'key' is missing")
    w.WriteHeader(400)
    return
  }

  busstopid := busStopID[0]
  // check the cache according to the busstopid 
  val, err := cache.Get(busstopid)
  if err != ttlcache.ErrNotFound{
    json.NewEncoder(w).Encode(val) 
    return
  }

  url := "http://dummy.uwave.sg/busstop/" + busstopid 
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
  var busStop responses.BusStop
  json.Unmarshal((data), &busStop)

  var busTimingResponses []responses.BusTimingResponse
  for i := 0; i<len(busStop.Forecast); i++ {
    var busTimingResponse responses.BusTimingResponse
    busTimingResponse.Estimated_arrival_time= busStop.Forecast[i].Forecast_seconds
    busTimingResponse.Vehicle_id = busStop.Forecast[i].Vehicle_id
    busTimingResponse.Route_name = busStop.Forecast[i].Route.Name
    busTimingResponses = append(busTimingResponses, busTimingResponse);
  }

  // cache the busTimingResponse
  cache.Set(busstopid, busTimingResponses);
  
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)    
  json.NewEncoder(w).Encode(busTimingResponses) // Writes to the localhost page
}
