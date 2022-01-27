package storage

import (
	"sync"

	"github.com/leviska/lurker/base"
)

type Subscribe struct {
	handlers map[string][]chan<- *base.Lyrics
	storage  base.Storager
	lock     sync.RWMutex
}

func NewSubscribe(storage base.Storager) *Subscribe {
	return &Subscribe{
		handlers: map[string][]chan<- *base.Lyrics{},
		storage:  storage,
	}
}

func (s *Subscribe) Load(song *base.Song) (*base.Lyrics, error) {
	return s.storage.Load(song)
}

func (s *Subscribe) Save(lyrics *base.Lyrics) error {
	if err := s.storage.Save(lyrics); err != nil {
		return err
	}

	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, h := range s.handlers[lyrics.Song.String()] {
		h <- lyrics
	}
	return nil
}

func (s *Subscribe) Subscribe(song *base.Song) <-chan *base.Lyrics {
	s.lock.Lock()
	defer s.lock.Unlock()

	ch := make(chan *base.Lyrics, 1)
	s.handlers[song.String()] = append(s.handlers[song.String()], ch)
	return ch
}
