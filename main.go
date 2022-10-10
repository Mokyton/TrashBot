package main

import (
	"flag"
	tgClient "github.com/Mokyton/TrashBot/clients/telegram"
	"github.com/Mokyton/TrashBot/events/telegram"
	"github.com/Mokyton/TrashBot/storage/files"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
		)
	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("toke is not specified")
	}
	return *token
}
