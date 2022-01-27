package base

import (
	"net/url"

	"github.com/leviska/lurker/format"
)

type Song struct {
	Artist string `json:"artist"`
	Name   string `json:"name"`
}

func NewSong(artist string, name string) Song {
	return Song{
		Artist: artist,
		Name:   name,
	}
}

func (s *Song) Format() Song {
	res := Song{}
	res.Name = format.FormatSongString(s.Name)
	res.Artist = format.FormatSongString(s.Artist)
	return res
}

func (s *Song) String() string {
	return s.Artist + " " + s.Name
}

type Lyrics struct {
	Song   Song    `json:"song"`
	Text   string  `json:"text"`
	Source url.URL `json:"source"`
}
