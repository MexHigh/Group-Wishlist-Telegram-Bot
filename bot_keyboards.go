package main

import (
	"git.leon.wtf/leon/group-wishlist-telegram-bot/wishlist"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// makeUsernameKeyboard creates an InlineKeyboardMarkup that can be directly
// assigned to msg.ReplyMarkup. Every User is listed in its own line.
func makeUsernameKeyboard(usernames ...wishlist.Username) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0)
	for _, username := range usernames {
		rows = append(rows,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("@"+string(username), string(username)),
			),
		)
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

// makeWishKeyboard creates an InlineKeyboardMarkup that can be directly
// assigned to msg.ReplyMarkup. Every Wish is listed in its own line.
func makeWishKeyboard(wishes ...wishlist.Wish) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0)
	for _, wish := range wishes {
		rows = append(rows,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(wish.Wish, wish.Wish),
			),
		)
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

// TODO find a way to destinguish callbacks from one another
