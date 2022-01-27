package parsers_test

import (
	"github.com/leviska/lurker/base"
	"github.com/leviska/lurker/crawler"
	"github.com/leviska/lurker/storage"
)

func DefaultEnv() (*storage.Subscribe, *crawler.Crawler) {
	store := storage.NewSubscribe(storage.NewMemory())
	return store, crawler.NewCrawler(store, crawler.NewClienter(crawler.NewRateLimiter(1)))
}

var Songs = []base.Song{
	base.NewSong("the black keys", "little black submarines"),
	base.NewSong("the kills", "gum"),
	base.NewSong("the constellation", "on my way up"),
	base.NewSong("The Kills", "Gypsy Death & You"),
}
