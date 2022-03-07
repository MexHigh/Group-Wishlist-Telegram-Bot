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

		if !update.Message.Chat.IsGroup() {
			bot.Send(tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"This bot can only be used in group chats",
			))
			continue
		}

		if update.CallbackQuery != nil {
			handleCallbackQuery(bot, update)
			continue
		}

		if update.Message != nil && update.Message.IsCommand() { // ignore any non-command Messages
			handleCommand(bot, update)
			continue
		}

	}

}

func handleCallbackQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	chatID := update.CallbackQuery.Message.Chat.ID
	//username := wishlist.Username(update.CallbackQuery.Message.From.UserName)
	msg := tgbotapi.NewMessage(chatID, "")
	msg.ParseMode = tgbotapi.ModeMarkdown

	wishlist, err := wishlist.GetWishlist(chatID, wishlist.Username(update.CallbackQuery.Data))
	if err != nil {
		msg.Text = err.Error()
	} else {
		msg.Text = fmt.Sprintf("*Wishlist for @%s*\n", update.CallbackQuery.Data)
		msg.Text += wishlist.String()
	}

	bot.Send(msg)
	bot.Send(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)) // this answers the callback query

}

func handleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	chatID := update.Message.Chat.ID
	username := wishlist.Username(update.Message.From.UserName)
	msg := tgbotapi.NewMessage(chatID, "")
	msg.ParseMode = tgbotapi.ModeMarkdown

	switch update.Message.Command() {
	case CommandStart:
		msg.Text = "Not implemented" // TODO
	case CommandHelp:
		msg.Text = "Got help" // TODO
	case CommandWish:
		if err := wishlist.AddWish(chatID, username, &wishlist.Wish{
			WishedAt: time.Now(),
			Wish:     update.Message.CommandArguments(),
		}); err != nil {
			msg.Text = err.Error()
		} else {
			msg.Text = "Wish created :)"
		}
	case CommandWishlist:
		users, err := wishlist.GetUsersWithWishes(chatID)
		if err != nil {
			msg.Text = err.Error()
		} else {
			msg.Text = "Which wishlist do you want to take a look at?" + "\n"
			msg.Text += "_(users that are not listed do not have any wishes)_"
			msg.ReplyMarkup = makeUsernameKeyboard(users...)
			// see handleCallbackQuery to continue with this conversation
		}
	case CommandFulfill:
		wishID, err := strconv.Atoi(update.Message.CommandArguments())
		if err != nil {
			msg.Text = "Argument must be an integer to the wish"
		} else {
			if err := wishlist.FulfillWish(chatID, username, wishID); err != nil {
				msg.Text = err.Error()
			} else {
				msg.Text = fmt.Sprintf("Wish %d fulfilled :)", wishID)
			}
		}
	}

	bot.Send(msg)

}
