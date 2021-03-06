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
// The callback data is composed like '/commandName/userID.username'
func makeUsernameKeyboard(commandName string, userInfos ...*wishlist.UserInfo) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0)
	for _, userInfo := range userInfos {
		cbData := fmt.Sprintf("/%s/%d.%s", commandName, userInfo.ID, userInfo.Username)
		rows = append(rows,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(userInfo.Username, cbData),
			),
		)
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

// makeWishlistKeyboard creates an InlineKeyboardMarkup that can be directly
// assigned to msg.ReplyMarkup. Every Wish is listed in its own line.
//
// The callback data is composed like '/commandName/userID.username.wishID'.
// Use `extractCallbackData` to get the
func makeWishlistKeyboard(commandName string, userInfo wishlist.UserInfo, skipFulfilled bool, wishlist *wishlist.Wishlist) tgbotapi.InlineKeyboardMarkup {
	rows := make([][]tgbotapi.InlineKeyboardButton, 0)
	for realWishID, wish := range wishlist.Wishes {
		if skipFulfilled && wish.Fulfilled {
			continue
		}
		wishID := realWishID + 1
		text := fmt.Sprintf("%d. %s", wishID, wish.Wish)
		cbData := fmt.Sprintf("/%s/%d.%s.%d", commandName, userInfo.ID, userInfo.Username, wishID)
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
