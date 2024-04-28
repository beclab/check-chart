package models

type ResponseBase struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
	Data any    `json:"data,omitempty"`
}
