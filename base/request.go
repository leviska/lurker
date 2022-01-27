package base

import "net/url"

type Request struct {
	URL    *url.URL
	Parser Parser

	Song *Song // optional
}

func NewRequest(u *url.URL, p Parser) *Request {
	return &Request{URL: u, Parser: p}
}
