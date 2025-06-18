package contest

import (
	"fmt"
	"net/url"
	"strings"
)

type QuestionResourceType int

const (
	ResourceOfUndefined QuestionResourceType = iota
	ResourceOfHtml
)

type QuestionConfig struct {
	Contest      string
	QuestionTask string
	QuestionId   string
	ResourceType QuestionResourceType
}

func (c *QuestionConfig) toQuestion() *question {
	co := new(contest)
	co.name = c.Contest

	q := new(question)
	q.contest = co
	q.task = c.QuestionTask
	q.id = c.QuestionId

	return q
}

type question struct {
	contest *contest
	task    string
	id      string
}

type qUrlBuilder struct {
	_scheme string
	_host   string
	_paths  []string
}

func newQUrlBuilder() *qUrlBuilder {
	b := new(qUrlBuilder)
	return b
}

func (b *qUrlBuilder) scheme(s string) *qUrlBuilder {
	b._scheme = s
	return b
}

func (b *qUrlBuilder) host(s string) *qUrlBuilder {
	b._host = s
	return b
}

func (b *qUrlBuilder) addPath(p string) *qUrlBuilder {
	b._paths = append(b._paths, p)
	return b
}

func (b *qUrlBuilder) build() *url.URL {
	p := strings.Join(b._paths, "/")

	u := &url.URL{
		Scheme: b._scheme,
		Host:   b._host,
		Path:   p,
	}

	return u
}

func (q *question) toQuestionUrl() *url.URL {
	builder := newQUrlBuilder()

	builder.scheme(urlScheme)
	builder.host(urlHost)
	builder.addPath(urlPathContests)
	builder.addPath(q.contest.name)
	builder.addPath(urlPathTasks)
	builder.addPath(fmt.Sprintf("%s_%s", q.task, q.id))

	return builder.build()
}
