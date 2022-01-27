package parsers_test

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/leviska/lurker/base"
	"github.com/leviska/lurker/parsers"
	"github.com/stretchr/testify/require"
)

func TestGenius(t *testing.T) {
	store, crawl := DefaultEnv()
	song := &Songs[0]
	done := store.Subscribe(song)

	genius := &parsers.Genius{}
	u, err := url.Parse("https://genius.com/the-black-keys-little-black-submarines-lyrics")
	require.NoError(t, err)
	req := base.NewRequest(u, genius)
	req.Song = song
	crawl.Request(req)

	fmt.Println(<-done)
	crawl.Stop()
}
