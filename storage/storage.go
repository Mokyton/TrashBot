package storage

import (
	"crypto/sha1"
	"fmt"
	"github.com/Mokyton/TrashBot/lib/e"
	"io"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL      string
	UserName string
	//Created time.Time
}

func (p Page) Hash() (string, error) {
	h := sha1.New()
	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("cant calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
