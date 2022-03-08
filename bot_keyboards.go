package main

import (
	"fmt"
	"strings"

	"git.leon.wtf/leon/group-wishlist-telegram-bot/wishlist"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// makeUsernameKeyboard creates an InlineKeyboardMarkup that can be directly
// assigned to msg.ReplyMarkup. Every User is listed in its own line.
//
// The callback data is composed like '/commandName/username'
func makeUsernameKeyboard(commandName string, usernames ...wishlist.Username) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0)
	for _, username := range usernames {
		cbData := fmt.Sprintf("/%s/%s", commandName, string(username))
		rows = append(rows,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("@"+string(username), cbData),
			),
		)
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

// makeWishlistKeyboard creates an InlineKeyboardMarkup that can be directly
// assigned to msg.ReplyMarkup. Every Wish is listed in its own line.
//
// The callback data is composed like '/commandName/username.wishID'.
// Use `extractCallbackData` to get the
func makeWishlistKeyboard(commandName string, username wishlist.Username, wishes wishlist.Wishlist) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0)
	for realWishID, wish := range wishes {
		wishID := realWishID + 1
		text := fmt.Sprintf("%d. %s", wishID, wish.Wish)
		cbData := fmt.Sprintf("/%s/%s.%d", commandName, string(username), wishID)
		rows = append(rows,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(text, cbData),
			),
		)
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

// extractCallbackData returns the command name the callback data
// was returned for and the callback data itself
func extractCallbackData(cbData string) (string, string) {
	s := strings.Split(cbData, "/")
	return s[1], s[2]
}
