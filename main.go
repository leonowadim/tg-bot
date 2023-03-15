package main

import (
	"context"
	"flag"
	"log"
	tgClient "telegramBot/clients/telegram"
	"telegramBot/consumer/event-consumer"
	"telegramBot/events/telegram"
	"telegramBot/storage/sqlite"
)

const (
	sqliteStoragePath = "data/sqlite/storage.db"
	batchSize         = 100
	tgBotHost         = "api.telegram.org" // СДЕЛАТЬ ФЛАГОМ !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
)

func main() {
	//s := files.New(storagePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatal("can not connect to storage: ", err)
	}
	if err := s.Init(context.TODO()); err != nil {
		log.Fatal("can not init storage: ", err)
	}
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
	)

	log.Print("service started")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal()
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		" ",
		"token for access to tg bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}

//./telegramBot -tg-bot-token '5791173406:AAHxmWZzAnNXCsTfThp9LzzZ4VOch-oZIQQ'
