package parsers_test

import (
	"testing"
	"time"

	"github.com/leviska/lurker/parsers"
	"github.com/stretchr/testify/assert"
)

func TestGoogle(t *testing.T) {
	store, crawl := DefaultEnv()

	for _, song := range Songs {
		notify := store.Subscribe(&song)
		u := parsers.NewGoogleRequest(&song)
		crawl.Request(u)
		lyrics := <-notify
		assert.Greater(t, len(lyrics.Text), 200)
		time.Sleep(time.Second)
	}

	crawl.Stop()
}
