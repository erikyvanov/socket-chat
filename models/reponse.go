package models

type Response struct {
	Ok     bool        `json:"ok"`
	Errors interface{} `json:"errors,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
