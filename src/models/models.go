package models

type RequestBody struct {
	Evidence map[string]int `json:"evidence"`
	Target   string         `json:"target"`
}

type ResponseBody struct {
	Maladie string `json:"maladie"`
	Urgence string `json:"urgence"`
}
