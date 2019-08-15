package commands

import "strings"

type Response struct {
	Data []byte
}

func NewResponse(d ...string) *Response {
	return &Response{
		Data: []byte(strings.Join(d, " ")),
	}
}

func (r *Response) AddData(d string) {
	r.Data = append(r.Data, []byte(" "+d)...)
}

func (r *Response) SetBytes(d []byte) {
	r.Data = append(d, '\n')
}

func (r *Response) Bytes() []byte {
	return r.Data
}
