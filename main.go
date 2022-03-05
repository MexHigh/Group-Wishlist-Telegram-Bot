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
	Config = confTemp // assign to global value

	bot, err := tgbotapi.NewBotAPI(Config.BotToken)
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	log.Println("Authorized on bot account @" + bot.Self.UserName)

	log.Println("Setting bot commands")
	if _, err := addBotCommands(bot); err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)
	for update := range updates {

		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		switch update.Message.Command() {
		case CommandStart:
			if !update.Message.Chat.IsGroup() {
				msg.Text = "This bot can only be used in group chats. I'm sorry :("
			} else {
				msg.Text = "Ok, I'm in a group"
			}
		case CommandHelp:
			msg.Text = "Got help"
		case CommandWish:
			msg.Text = "Got wish"
		case CommandWishlist:
			msg.Text = "Got wishlist"
		case CommandFulfilled:
			msg.Text = "Got fulfilled"
		}

		bot.Send(msg)

	}

}
