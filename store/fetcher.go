package store

import (
	"io/ioutil"
	"net/http"
)

type Response struct {
	Body []byte
	Err  error
}

//TODO implementation with channels
type Fetcher struct {
}

func (f Fetcher) Fetch(in []string) (out []Response) {
	for _, url := range in {
		out = append(out, f.Request(url))
	}
	return out
}

func (f Fetcher) Request(url string) Response {
	resp, err := http.Get(url)
	if err != nil {
		return Response{Err: err}
	}
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{Err: err}
	}
	return Response{Body: raw}
}
