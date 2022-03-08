package main

import (
	"fmt"
	"strconv"
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

// makeWishKeyboard creates an InlineKeyboardMarkup that can be directly
// assigned to msg.ReplyMarkup. Every Wish is listed in its own line.
//
// The callback data is composed like '/commandName/wishID'.
// Use `extractCallbackData` to get the
func makeWishKeyboard(commandName string, wishes ...wishlist.Wish) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0)
	for wishID, wish := range wishes {
		cbData := fmt.Sprintf("/%s/%d", commandName, wishID)
		rows = append(rows,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(wishID)+wish.Wish, cbData),
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

// matchCallbackCommandName returns only the command name
// for which callback data was returned.
//
// This can be useful for switch case statements.
func matchCallbackCommandName(cbData string) string {
	command, _ := extractCallbackData(cbData)
	return command
}
