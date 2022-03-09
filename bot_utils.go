package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// firstAccountInfoAvailable gets the first valid string
// identifier for a Telegram account in the following order:
//
// full name > first name only > username (with @)
//
// If none of them are available, "unknown user" + account ID is returned
func firstAccountInfoAvailable(user *tgbotapi.User) (identifier string) {
	if user.FirstName != "" {
		identifier = user.FirstName
		if user.LastName != "" {
			identifier += " " + user.LastName
		}
	} else if user.UserName != "" {
		identifier = fmt.Sprintf("@%s", user.UserName)
	} else {
		identifier = fmt.Sprintf("unknown user (%d)", user.ID)
	}
	return
}
