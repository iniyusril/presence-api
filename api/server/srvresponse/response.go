package srvresponse

type Response struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	ErrorCode int         `json:"error_code,omitempty"`
	Data      interface{} `json:"data"`
}
