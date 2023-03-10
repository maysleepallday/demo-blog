package handler

type ResponseData struct {
	Code int         `json:"code" example:"200"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}
