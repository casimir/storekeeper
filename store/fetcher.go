package store

import (
	"io/ioutil"
	"net/http"
	"sync"
)

type Response struct {
	Body []byte
	Err  error
}

type Fetcher struct {
}

func (f Fetcher) cFetch(in []string, nb int) (out []Response) {
	sem := make(chan bool, nb)
	for i := 0; i < nb; i++ {
		sem <- true
	}

	outCh := make(chan Response)
	var wg sync.WaitGroup
	for _, url := range in {
		wg.Add(1)
		go func() {
			<-sem
			outCh <- f.Request(url)
			sem <- true
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(outCh)
	}()

	for r := range outCh {
		out = append(out, r)
	}
	return out
}

func (f Fetcher) Fetch(in []string) []Response {
	return f.cFetch(in, http.DefaultMaxIdleConnsPerHost)
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
