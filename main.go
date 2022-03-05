package main

import (
	"flag"
	"log"

	"git.leon.wtf/leon/group-wishlist-telegram-bot/files"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	configPath *string       = flag.String("config", "./config.json", "Path to the config.json file")
	Config     *files.Config = nil
)

func main() {

	flag.Parse()

	log.Println("Loading config from", *configPath)
	confTemp, err := files.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}
	Config = confTemp

	bot, err := tgbotapi.NewBotAPI(Config.BotToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Println("Authorized on bot account @" + bot.Self.UserName)

}
