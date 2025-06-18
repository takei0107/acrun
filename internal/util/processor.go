package util

import "strings"

type Replacer struct {
	replacements map[string]string
}

func NewReplacer() *Replacer {
	r := new(Replacer)
	r.replacements = make(map[string]string)
	return r
}

func (r *Replacer) AddReplacements(p string, v string) {
	r.replacements[p] = v
}

func (r *Replacer) ReplaceStr(s string) string {
	for p, v := range r.replacements {
		s = strings.ReplaceAll(s, p, v)
	}
	return s
}

func (r *Replacer) ReplaceStrSlice(ss []string) []string {
	n := make([]string, len(ss), len(ss))
	for i, s := range ss {
		s = r.ReplaceStr(s)
		n[i] = s
	}
	return n
}
