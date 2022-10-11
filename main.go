package main

import (
	"flag"
	tgClient "github.com/Mokyton/TrashBot/clients/telegram"
	event_consumer "github.com/Mokyton/TrashBot/consumer/event-consumer"
	"github.com/Mokyton/TrashBot/events/telegram"
	"github.com/Mokyton/TrashBot/storage/files"
	"log"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "rl_storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("toke is not specified")
	}
	return *token
}
