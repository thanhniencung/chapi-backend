package model

type Response struct {
	StatusCode int `json:"statusCode,omitempty"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
}