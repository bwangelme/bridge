package httperr

import "fmt"

type HTTPErr struct {
	Err  error
	Code int
}

func New400Error(err error) *HTTPErr {
	return NewHTTPErr(400, err)
}

func New500Error(err error) *HTTPErr {
	return NewHTTPErr(500, err)
}

func NewHTTPErr(code int, err error) *HTTPErr {
	return &HTTPErr{
		Err:  err,
		Code: code,
	}
}

func (h *HTTPErr) Error() string {
	return fmt.Sprintf("[%d] %v", h.Code, h.Err)
}

func (h *HTTPErr) Unwrap() error {
	return h.Err
}
