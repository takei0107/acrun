package contest

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/takei0107/acrun/internal/util"
)

const (
	urlScheme       = "https"
	urlHost         = "atcoder.jp"
	urlPathContests = "contests"
	urlPathTasks    = "tasks"
)

type Inputs []string
type Outputs []string

type SampleInOuts struct {
	No      int
	Inputs  Inputs
	Outputs Outputs
}

type ResourceParser interface {
	ParseSample() ([]*SampleInOuts, error)
}

func getResourceParser(t QuestionResourceType, r io.Reader) (ResourceParser, error) {
	var rp ResourceParser
	var err error
	switch t {
	case ResourceOfHtml:
		p := new(htmlParser)
		p.reader = r
		rp = p
		err = nil
	case ResourceOfUndefined:
		rp = nil
		err = fmt.Errorf("[acrun] invalid quesion resource type=%d\n", t)
	}

	return rp, err
}

func httpClosure(u *url.URL) func(*http.Response) error {
	return func(r *http.Response) error {
		if r.StatusCode != 200 {
			return fmt.Errorf("[acrun] Get Request with unexpected code=%d, url=%s\n", r.StatusCode, u.String())
		}
		return nil
	}
}

func getReader(u *url.URL) (io.ReadCloser, error) {
	fmt.Printf("[acrun] fetch file from %s\n", u.String())
	r, err := util.Get(u, httpClosure(u))
	if err != nil {
		return nil, err
	}
	fmt.Printf("[acrun] ok\n")

	return r, nil
}

func GetSampleInOutsSlice(c *QuestionConfig) ([]*SampleInOuts, error) {
	q := c.toQuestion()
	u := q.toQuestionUrl()

	r, err := getReader(u)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	p, err := getResourceParser(c.ResourceType, r)
	if err != nil {
		return nil, err
	}

	return p.ParseSample()
}
