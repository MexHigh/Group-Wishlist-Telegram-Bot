package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	CommandStart     string = "start"
	CommandHelp      string = "help"
	CommandWish      string = "wish"
	CommandWishlist  string = "wishlist"
	CommandFulfilled string = "fulfilled"
)

func addBotCommands(bot *tgbotapi.BotAPI) (*tgbotapi.APIResponse, error) {
	return bot.Request(tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     CommandStart,
			Description: "Initializes bot",
		},
		tgbotapi.BotCommand{
			Command:     CommandHelp,
			Description: "Shows help message",
		},
		tgbotapi.BotCommand{
			Command:     CommandWish,
			Description: "Adds a new wish",
		},
		tgbotapi.BotCommand{
			Command:     CommandWishlist,
			Description: "Shows wishes of someone",
		},
		tgbotapi.BotCommand{ // TODO don't know how to do this. by ID or string search or inline keyboard?
			Command:     CommandFulfilled,
			Description: "Fulfills someones wish",
		},
	))
}
