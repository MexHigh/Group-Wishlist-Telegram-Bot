package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
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
	log.Println("Authorized on bot account @" + bot.Self.UserName)
	//bot.Debug = true

	log.Println("Setting bot commands")
	if _, err := addBotCommands(bot); err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30
	updates := bot.GetUpdatesChan(u)
	for update := range updates {

		if update.Message != nil && !update.Message.Chat.IsGroup() {
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
	commandName, cbDataPayload := extractCallbackData(update.CallbackData())
	msg := tgbotapi.NewMessage(chatID, "")
	msg.ParseMode = tgbotapi.ModeMarkdown

	switch commandName {
	case CommandWishlist: // callback payload format: username
		wishlist, err := wishlist.GetWishlist(chatID, wishlist.Username(cbDataPayload))
		if err != nil {
			msg.Text = beautifulError(err)
		} else {
			msg.Text = fmt.Sprintf("*Wishlist for @%s*", cbDataPayload) + "\n"
			msg.Text += wishlist.String()
		}
	case CommandFulfill: // callback payload format: username.wishID
		split := strings.Split(cbDataPayload, ".")
		username := wishlist.Username(split[0])
		wishID, err := strconv.Atoi(split[1])
		if err != nil {
			msg.Text = beautifulError(err)
			break
		}
		if err := wishlist.FulfillWish(chatID, username, wishID); err != nil {
			msg.Text = beautifulError(err)
		} else {
			msg.Text = fmt.Sprintf("Wish %d marked as fulfilled", wishID)
		}
	}

	// remove markup (inline keyboard) (unused as it does not make any sense to only remove the keyboard)
	/*editMarkup := tgbotapi.NewEditMessageReplyMarkup(chatID, update.CallbackQuery.Message.MessageID, tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: make([][]tgbotapi.InlineKeyboardButton, 0),
	})
	bot.Send(editMarkup)*/

	bot.Send(msg)
	bot.Send(tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)) // this answers the callback query

}

func handleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {

	chatID := update.Message.Chat.ID
	username := wishlist.Username(update.Message.From.UserName)
	msg := tgbotapi.NewMessage(chatID, "")
	msg.ParseMode = tgbotapi.ModeMarkdown

	switch update.Message.Command() {
	case CommandHelp:
		msg.Text = "Not implemented yet" // TODO
	case CommandWish:
		args := update.Message.CommandArguments()
		if args == "" {
			msg.Text = "Please provide your wish with your command!" + "\n"
			msg.Text += "Example: `/wish Diamond necklace`"
			break
		}
		if err := wishlist.AddWish(chatID, username, &wishlist.Wish{
			WishedAt: time.Now(),
			Wish:     args,
		}); err != nil {
			msg.Text = beautifulError(err)
		} else {
			msg.Text = "Wish created"
		}
	case CommandWishlist:
		users, err := wishlist.GetUsersWithWishes(chatID)
		if err != nil {
			msg.Text = beautifulError(err)
		} else {
			msg.Text = "Which wishlist do you want to take a look at?" + "\n"
			msg.Text += "_(users that are not listed do not have any wishes)_"
			msg.ReplyMarkup = makeUsernameKeyboard(CommandWishlist, users...)
			// see handleCallbackQuery to continue with this conversation
		}
	case CommandFulfill:
		list, err := wishlist.GetWishlist(chatID, username)
		if err != nil {
			msg.Text = beautifulError(err)
		} else {
			msg.Text = "Which wish of yours do you want to mark as fulfilled?"
			msg.ReplyMarkup = makeWishlistKeyboard(CommandFulfill, username, list)
			// see handleCallbackQuery to continue with this conversation
		}
	}

	bot.Send(msg)

}
