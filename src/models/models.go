package models

type RequestBody struct {
	Evidence map[string]int `json:"evidence"` 
	Target   string         `json:"target"`   
}

type ResponseBody struct {
	Target      string             `json:"target"`
	Predictions map[string]float64 `json:"predictions"` 
}