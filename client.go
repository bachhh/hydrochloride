package hydrochloride

import (
	"net/http"
)

// HClient a boring http client
type HClient struct {
	stdClient *http.Client

	// implement for most cookie jar
	roundtrip     http.RoundTripper
	cookieJar     http.CookieJar
	checkRedirect func(req *http.Request, via []*http.Request) error
}

func NewClient(base *http.Client, options ...ClientOption) (hcl *HClient, err error) {
	if base != nil {
		hcl = &HClient{stdClient: base}
	}
	hcl = &HClient{
		stdClient: &http.Client{},
	}

	for i := range options {
		options[i](hcl)
	}
	return hcl, nil
}

type ClientOption func(*HClient)

// WithRoundTripper
func WithRoundTripper(roundtrip http.RoundTripper) ClientOption {
	return func(hcl *HClient) {
		hcl.roundtrip = roundtrip
	}
}

func WithCookieJar(cookieJar http.CookieJar) ClientOption {
	return func(hcl *HClient) {
		hcl.cookieJar = cookieJar
	}
}

func WithRedirect(checkRedirect func(req *http.Request, via []*http.Request) error) ClientOption {
	return func(hcl *HClient) {
		hcl.checkRedirect = checkRedirect
	}
}

// // TODO
// func (c *HClient) CloseIdleConnections()
// func (c *HClient) Do(req *Request) (*Response, error)
// func (c *HClient) Get(url string) (resp *Response, err error)
// func (c *HClient) Head(url string) (resp *Response, err error)
// func (c *HClient) Post(url, contentType string, body io.Reader) (resp *Response, err error)
// func (c *HClient) PostForm(url string, data url.Values) (resp *Response, err error)
