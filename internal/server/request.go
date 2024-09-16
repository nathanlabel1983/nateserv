package server

import "strings"

type Request struct {
	Method      string
	Path        string
	HttpVersion string

	Headers map[string]string
}

func newRequest(data []string) Request {
	reqParts := strings.Split(data[0], " ")
	r := Request{
		Method:      reqParts[0],
		Path:        reqParts[1],
		HttpVersion: reqParts[2],
		Headers:     make(map[string]string),
	}

	for _, v := range data[1:] {
		if len(v) == 0 {
			continue
		}
		nv := strings.Split(v, ": ")
		if len(nv) != 2 {
			continue
		}
		r.Headers[nv[0]] = nv[1]
	}
	return r
}
