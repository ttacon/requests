package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type RequestBuilder interface {
	Cookie(c *http.Cookie) RequestBuilder
	BasicAuth(username, password string) RequestBuilder
	AddHeader(key, value string) RequestBuilder
	SetHeader(key, value string) RequestBuilder

	Params(m map[string][]string) RequestBuilder
	Body(r io.Reader) RequestBuilder
	JSONBody(i interface{}) RequestBuilder

	Request() (*http.Request, error)

	Do() (*http.Response, error)
	DoWithClient(client *http.Client) (*http.Response, error)
}

type requestBuilder struct {
	req *http.Request
	err error
}

func Get(url string) RequestBuilder {
	req, err := http.NewRequest("GET", url, nil)
	return &requestBuilder{
		req: req,
		err: err,
	}
}

func Post(url string) RequestBuilder {
	req, err := http.NewRequest("POST", url, nil)
	return &requestBuilder{
		req: req,
		err: err,
	}
}

func Put(url string) RequestBuilder {
	req, err := http.NewRequest("PUT", url, nil)
	return &requestBuilder{
		req: req,
		err: err,
	}
}

func Delete(url string) RequestBuilder {
	req, err := http.NewRequest("DELETE", url, nil)
	return &requestBuilder{
		req: req,
		err: err,
	}
}

func Head(url string) RequestBuilder {
	req, err := http.NewRequest("HEAD", url, nil)
	return &requestBuilder{
		req: req,
		err: err,
	}
}

func (r *requestBuilder) Params(vals map[string][]string) RequestBuilder {
	query := r.req.URL.Query()
	for k, vlist := range vals {
		for _, v := range vlist {
			query.Add(k, v)
		}
	}
	r.req.URL.RawQuery = query.Encode()
	return r
}

func (r *requestBuilder) Body(rc io.Reader) RequestBuilder {
	r.req.Body = ioutil.NopCloser(rc)
	return r
}

func (r *requestBuilder) JSONBody(i interface{}) RequestBuilder {
	data, err := json.Marshal(i)
	if r.err == nil && err != nil {
		r.err = err
	}
	r.req.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	r.req.Header.Add("Content-Type", "application/json")
	r.req.ContentLength = int64(len(data))
	return r
}

func (r *requestBuilder) Cookie(c *http.Cookie) RequestBuilder {
	r.req.AddCookie(c)
	return r
}

func (r *requestBuilder) BasicAuth(username, password string) RequestBuilder {
	r.req.SetBasicAuth(username, password)
	return r
}

func (r *requestBuilder) AddHeader(key, value string) RequestBuilder {
	r.req.Header.Add(key, value)
	return r
}

func (r *requestBuilder) SetHeader(key, value string) RequestBuilder {
	r.req.Header.Set(key, value)
	return r
}

func (r *requestBuilder) Request() (*http.Request, error) {
	return r.req, r.err
}

func (r *requestBuilder) Do() (*http.Response, error) {
	if r.err != nil {
		return nil, r.err
	}
	return http.DefaultClient.Do(r.req)
}

func (r *requestBuilder) DoWithClient(client *http.Client) (*http.Response, error) {
	if r.err != nil {
		return nil, r.err
	}
	return client.Do(r.req)
}
