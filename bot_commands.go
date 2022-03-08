package main

import (
	t "git.leon.wtf/leon/group-wishlist-telegram-bot/translator"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	CommandHelp     string
	CommandWish     string
	CommandWishlist string
	CommandFulfill  string
)

func addBotCommands(bot *tgbotapi.BotAPI) (*tgbotapi.APIResponse, error) {
	// initialitze the constants
	CommandHelp = t.G("command_help")
	CommandWish = t.G("command_wish")
	CommandWishlist = t.G("command_wishlist")
	CommandFulfill = t.G("command_fulfill")

	return bot.Request(tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{
			Command:     CommandHelp,
			Description: t.G("command_help_desc"),
		},
		tgbotapi.BotCommand{
			Command:     CommandWish,
			Description: t.G("command_wish_desc"),
		},
		tgbotapi.BotCommand{
			Command:     CommandWishlist,
			Description: t.G("command_wishlist_desc"),
		},
		tgbotapi.BotCommand{
			Command:     CommandFulfill,
			Description: t.G("command_fulfill_desc"),
		},
	))
}
