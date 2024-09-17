package server

import "fmt"

type Response struct {
	Version string
	Status  string
	Reason  string

	Headers map[string]string
	Body    string
}

const (
	DefaultHTTPVersion = "HTTP/1.1"

	StatusOK       = "200"
	StatusOKReason = "OK"
)

func (r *Response) GetResponseString() string {
	responseString := fmt.Sprintf("%s %s %s\n", r.Version, r.Status, r.Reason)
	for k, v := range r.Headers {
		responseString += fmt.Sprintf("%s: %s\n", k, v)
	}
	responseString += "\r\n"
	responseString += r.Body
	return responseString
}

func NewResponse(status, reason, body string) Response {
	res := Response{
		Version: DefaultHTTPVersion,
		Status:  status,
		Reason:  reason,
		Headers: make(map[string]string),
		Body:    body,
	}
	return res
}
