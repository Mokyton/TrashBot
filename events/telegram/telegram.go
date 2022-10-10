package telegram

import (
	"github.com/Mokyton/TrashBot/clients/telegram"
	"github.com/Mokyton/TrashBot/events"
	"github.com/Mokyton/TrashBot/lib/e"
	"github.com/Mokyton/TrashBot/storage"
)

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.Storage
}

type Meta struct {
	ChatID   int
	Username string
}

func New(client *telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      nil,
		storage: nil,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	update, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get events", err)
	}
	res := make([]events.Event, 0, len(update))

	for _, u := range update {
		res = append(res, event(u))
	}
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}
	if updType == events.Message {
		res.Meta{

		}
	}
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}
	return events.Message
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text
}
