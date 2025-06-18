package contest

import (
	"io"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

var reI = regexp.MustCompile(`^Sample Input \d+$`)
var reIc = regexp.MustCompile(`^Sample Input (\d+)$`)
var reO = regexp.MustCompile(`^Sample Output \d+$`)
var reOc = regexp.MustCompile(`^Sample Output (\d+)$`)

type htmlParser struct {
	reader io.Reader
}

func nextText(z *html.Tokenizer) error {
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			return z.Err()
		case html.TextToken:
			return nil
		}
	}
}

func handleErrorToken(err error) error {
	if err == io.EOF {
		return nil
	} else {
		return err
	}
}

func getOrNewSampleInOuts(s []*SampleInOuts, n int) *SampleInOuts {
	for _, ss := range s {
		if ss.No == n {
			return ss
		}
	}
	r := new(SampleInOuts)
	r.No = n
	s = append(s, r)
	return r
}

func (p *htmlParser) ParseSample() ([]*SampleInOuts, error) {
	s := make([]*SampleInOuts, 10, 10)

	z := html.NewTokenizer(p.reader)

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			err := z.Err()
			err = handleErrorToken(err)
			if err != nil {
				return nil, err
			}
			return s, nil
		case html.TextToken:
			c := string(z.Text())

			if reI.MatchString(c) {
				ns := reIc.FindStringSubmatch(c)[1]
				n, err := strconv.Atoi(ns)
				if err != nil {
					return nil, err
				}

				err = nextText(z)
				if err != nil {
					err = handleErrorToken(err)
					if err != nil {
						return nil, err
					}
					return s, nil
				}

				c := string(z.Text())
				sp := strings.Split(c, "\r\n")

				ss := getOrNewSampleInOuts(s, n)

				for i, in := range sp {
					ss.Inputs[i] = in
				}

			} else if reO.MatchString(c) {
				ns := reOc.FindStringSubmatch(c)[1]
				n, err := strconv.Atoi(ns)
				if err != nil {
					return nil, err
				}

				err = nextText(z)
				if err != nil {
					err = handleErrorToken(err)
					if err != nil {
						return nil, err
					}
					return s, nil
				}

				c := string(z.Text())

				sp := strings.Split(c, "\r\n")

				ss := getOrNewSampleInOuts(s, n)

				for i, in := range sp {
					ss.Outputs[i] = in
				}
			}
		}
	}
}
