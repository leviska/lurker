package storage

import (
	"fmt"
	"sync"

	"github.com/leviska/lurker/base"
)

type Memory struct {
	data map[string]*base.Lyrics
	lock sync.RWMutex
}

func NewMemory() *Memory {
	return &Memory{
		data: map[string]*base.Lyrics{},
	}
}

func (s *Memory) Load(song *base.Song) (*base.Lyrics, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if lyrics, has := s.data[song.String()]; has {
		return lyrics, nil
	}
	return nil, &StorageErr{fmt.Errorf("not found")}
}

func (s *Memory) Save(lyrics *base.Lyrics) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data[lyrics.Song.String()] = lyrics
	return nil
}
