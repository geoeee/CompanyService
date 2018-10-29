package utils

import (
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

// RoundTripper ...
type RoundTripper struct {
	original http.RoundTripper
	h        http.Header
}

// NewRoundTripper ...
func NewRoundTripper(original http.RoundTripper, h http.Header) *RoundTripper {
	rt := new(RoundTripper)
	rt.original = original
	rt.h = h
	return rt
}

// RoundTrip ...
func (t *RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {

	if t.h["X-Request-Platform"] != nil {
		req.Header.Set("X-Request-Platform", fmt.Sprintf("%s", t.h["X-Request-Platform"][0]))
	}
	fmt.Println(spew.Sdump(t.h))
	resp, err := t.original.RoundTrip(req)
	if err == nil {

	} else {
		fmt.Println(err.Error())

		return nil, err

	}

	return resp, err
}
