package structs

type BusLine struct {
  BusLineName string `json:"name"`
  Vehicles AllVehicles  `json:"vehicles"`
}
type Vehicle struct {
  Bearing int `json:"bearing"` 
  Latitude string `json:"lat"`
  Longtitude string `json:"lon"` 
  Vehicle_id int `json:"vehicle_id"`
}

type BusLocationResponse struct{
  BusLineName string `json:"busline_name"` 
  Latitude string `json:"latitude"`
  Longtitude string `json:"longtitude"`
  Vehicle_id int `json:"vehicle_id"`
}

type AllVehicles []Vehicle
