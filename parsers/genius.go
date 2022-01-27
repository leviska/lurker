package parsers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/leviska/lurker/base"
	"github.com/leviska/lurker/format"
)

type Genius struct {
}

func (p *Genius) Parse(req *base.Request, doc *goquery.Document, requester base.Requester, saver base.Saver) error {
	builder := strings.Builder{}
	doc.Find("div[data-lyrics-container=true]").Each(func(i int, s *goquery.Selection) {
		s.Contents().Each(func(i int, s *goquery.Selection) {
			if s.Is("br") {
				builder.WriteByte('\n')
			} else if s.Is("a") {
				html, _ := s.Html()
				s.SetHtml(format.AddBRNewLine(html))
				builder.WriteString(s.Text())
			} else if goquery.NodeName(s) == "#text" {
				builder.WriteString(s.Text())
			}
		})
	})
	lyrics := &base.Lyrics{
		Song:   *req.Song,
		Text:   format.FormatNewLine(builder.String()),
		Source: *req.URL,
	}
	saver.Save(lyrics)
	return nil
}
