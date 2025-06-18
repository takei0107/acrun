package util

import (
	"bufio"
	//"fmt"
	"io"
	"net/http"
	"net/url"
)

type FetchType int

func Get(u *url.URL, fn func(*http.Response) error) (io.ReadCloser, error) {
	r, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	err = fn(r)
	if err != nil {
		return nil, err
	}

	return r.Body, nil
}

func ReadToOuts(r io.Reader, outs *[]string) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		t := scanner.Text()
		//fmt.Printf("[acrun] debug: %s\n", t)
		*outs = append(*outs, t)
	}
}
