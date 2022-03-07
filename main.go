package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"git.leon.wtf/leon/group-wishlist-telegram-bot/config"
	"git.leon.wtf/leon/group-wishlist-telegram-bot/wishlist"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	configPath *string        = flag.String("config", "./config.json", "Path to the config.json file")
	Config     *config.Config = nil
)

func main() {

	flag.Parse()

	log.Println("Loading config from", *configPath)
	confTemp, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}
	Config = confTemp // assign to global value

	bot, err := tgbotapi.NewBotAPI(Config.BotToken)
	if err != nil {
		panic(err)
	}
	//bot.Debug = true

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

		chatID := update.Message.Chat.ID
		username := wishlist.Username(update.Message.From.UserName)
		msg := ""

		switch update.Message.Command() {
		case CommandStart:
			if !update.Message.Chat.IsGroup() {
				msg = "This bot can only be used in group chats. I'm sorry :("
			} else {
				msg = "Ok, I'm in a group"
			}
		case CommandHelp:
			msg = "Got help"
		case CommandWish:
			if err := wishlist.AddWish(chatID, username, &wishlist.Wish{
				WishedAt: time.Now(),
				Wish:     update.Message.CommandArguments(),
			}); err != nil {
				msg = err.Error()
			} else {
				msg = "Wish created :)"
			}
		case CommandWishlist:
			// TODO add mechanism to show inline keyboard for whom to show the wishlist for
			wishlist, err := wishlist.GetWishlist(chatID, username)
			if err != nil {
				msg = err.Error()
			} else {
				msg = fmt.Sprintf("*Wishlist for @%s*\n", username)
				msg += wishlist.String()
			}
		case CommandFulfill:
			wishID, err := strconv.Atoi(update.Message.CommandArguments())
			if err != nil {
				msg = "Argument must be an integer to the wish"
			} else {
				if err := wishlist.FulfillWish(chatID, username, wishID); err != nil {
					msg = err.Error()
				} else {
					msg = fmt.Sprintf("Wish %d fulfilled :)", wishID)
				}
			}
		}

		mdMsg := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
		mdMsg.ParseMode = tgbotapi.ModeMarkdown
		bot.Send(mdMsg)

	}

}
