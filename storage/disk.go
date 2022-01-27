package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/leviska/lurker/base"
)

type DiskFormat int

const (
	DiskJSON DiskFormat = iota
	DiskText
)

type Disk struct {
	folder string
	format DiskFormat
}

func NewDisk(folder string, format DiskFormat) *Disk {
	return &Disk{
		folder: folder,
		format: format,
	}
}

func (s *Disk) songFileName(song *base.Song) string {
	var ext string
	switch s.format {
	case DiskJSON:
		ext = "json"
	case DiskText:
		ext = "txt"
	}
	return filepath.Join(s.folder, song.String()) + "." + ext
}

func (s *Disk) loadJSON(file *os.File) (*base.Lyrics, error) {
	lyrics := &base.Lyrics{}
	err := json.NewDecoder(file).Decode(lyrics)
	if err != nil {
		return nil, err
	}
	return lyrics, nil
}

func (s *Disk) loadText(file *os.File) (*base.Lyrics, error) {
	text, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return &base.Lyrics{
		Text: string(text),
	}, nil
}

func (s *Disk) Load(song *base.Song) (*base.Lyrics, error) {
	file, err := os.Open(s.songFileName(song))
	if err != nil {
		return nil, &StorageErr{internal: err}
	}
	switch s.format {
	case DiskJSON:
		return s.loadJSON(file)
	case DiskText:
		return s.loadText(file)
	default:
		return nil, fmt.Errorf("wrong format")
	}
}

func (s *Disk) saveJSON(file *os.File, lyrics *base.Lyrics) error {
	return json.NewEncoder(file).Encode(lyrics)
}

func (s *Disk) saveText(file *os.File, lyrics *base.Lyrics) error {
	data := []byte(lyrics.Text)
	_, err := file.Write(data)
	return err
}

func (s *Disk) Save(lyrics *base.Lyrics) error {
	file, err := os.Create(s.songFileName(&lyrics.Song))
	if err != nil {
		return err
	}
	switch s.format {
	case DiskJSON:
		return s.saveJSON(file, lyrics)
	case DiskText:
		return s.saveText(file, lyrics)
	default:
		return fmt.Errorf("wrong format")
	}
}
