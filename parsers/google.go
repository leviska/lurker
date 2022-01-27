package parsers

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/leviska/lurker/base"
)

type Google struct {
}

func (p *Google) Parse(req *base.Request, doc *goquery.Document, requester base.Requester, saver base.Saver) error {
	doc.Find("div.g").Each(func(i int, s *goquery.Selection) {
		linkHref, _ := s.Find("a").Attr("href")
		linkText := strings.TrimSpace(linkHref)
		link, err := url.Parse(linkText)
		if err != nil {
			panic(err)
		}
		if parser, has := ParserMap[link.Host]; has {
			requester.Request(&base.Request{
				URL:    link,
				Parser: parser,
				Song:   req.Song,
			})
		}
	})
	return nil
}

func NewGoogleRequest(song *base.Song) *base.Request {
	str := strings.ReplaceAll(song.String(), " ", "+")
	link, err := url.Parse(fmt.Sprintf("https://www.google.com/search?q=%s+lyrics", str))
	if err != nil {
		panic(err)
	}
	return &base.Request{
		URL:    link,
		Parser: &Google{},

		Song: song,
	}
}
