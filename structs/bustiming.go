package structs

type BusStop struct {
  External_id string `json:"external_id"`
  Forecast AllForecast `json:"forecast"`
  Geometry AllGeometry `json:"geometry"`
  Id int `json:"id"`
  Name string `json:"name"`
  Name_en string `json:"name_en"`
  Name_ru string `json:"name_ru"`
  Nameslug string `json:"nameslug"`
  Resource_URI string `json:"resource_uri"`
}

type Route struct {
  Id int `json:"id"`
  Name string `json:"name"`
  Shortname string `json:"short_name"`
}

type Forecast struct {
  Forecast_seconds float32 `json:"forecast_seconds"`
  Route Route `json:"route"`
  Rv_id int `json:"rv_id"`
  Total_pass float32 `json:"total_pass"`
  Vehicle string `json:"vehicle"`
  Vehicle_id float32 `json:"vehicle_id`
}

type Geometry struct {
  External_id string `json:"external_id"`
  Lat string `json:"lat"`
  Lon string `json:"lon"`
  Seq int `json:"seq"`
}

type BusTimingResponse struct {
  Estimated_arrival_time float32 `json:"estimated_arrival_time"`
  Route_name string `json:"route_name"`
  Vehicle_id float32 `json:"vehicle_id"`
} 

type AllGeometry []Geometry
type AllForecast []Forecast


