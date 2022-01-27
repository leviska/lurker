package base

import (
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type Saver interface {
	Save(*Lyrics) error
}

type Loader interface {
	Load(song *Song) (*Lyrics, error)
}

type Storager interface {
	Saver
	Loader
}

type Requester interface {
	Request(*Request)
}

type Clienter interface {
	Get(*url.URL) *http.Request
	Put(*url.URL)
}

type Parser interface {
	Parse(*Request, *goquery.Document, Requester, Saver) error
}
