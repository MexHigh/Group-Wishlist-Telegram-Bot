package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
	"time"

	"git.leon.wtf/leon/group-wishlist-telegram-bot/config"
	t "git.leon.wtf/leon/group-wishlist-telegram-bot/translator"
	"git.leon.wtf/leon/group-wishlist-telegram-bot/wishlist"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	configPath *string = flag.String("config", "./config.json", "Path to the config.json file")
	language   *string = flag.String("language", "de", "Language of the bot ('en' or 'de')")
)

func init() {
	flag.Parse()
	t.SetLanguage(t.Language(*language))
}

func main() {

	log.Println("Loading config from", *configPath)
	conf, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(conf.BotToken)
	if err != nil {
		panic(err)
	}
	log.Println("Authorized on bot account @" + bot.Self.UserName)

	if conf.Debug {
		bot.Debug = true
	}

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
				t.G("This bot can only be used in group chats"),
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
			msg.Text = t.G("*Wishlist for @%s*", cbDataPayload) + "\n"
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
			msg.Text = t.G("Wish %d marked as fulfilled", wishID)
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
		msg.Text = t.G("Not implemented yet") // TODO
	case CommandWish:
		args := update.Message.CommandArguments()
		if args == "" {
			msg.Text = t.G("Please provide your wish with your command!") + "\n"
			msg.Text += t.G("Example: `/wish Diamond necklace`")
			break
		}
		if err := wishlist.AddWish(chatID, username, &wishlist.Wish{
			WishedAt: time.Now(),
			Wish:     args,
		}); err != nil {
			msg.Text = beautifulError(err)
		} else {
			msg.Text = t.G("Wish created")
		}
	case CommandWishlist:
		users, err := wishlist.GetUsersWithWishes(chatID)
		if err != nil {
			msg.Text = beautifulError(err)
		} else {
			msg.Text = t.G("Which wishlist do you want to take a look at?") + "\n"
			msg.Text += t.G("_(users that are not listed do not have any wishes)_")
			msg.ReplyMarkup = makeUsernameKeyboard(CommandWishlist, users...)
			// see handleCallbackQuery to continue with this conversation
		}
	case CommandFulfill:
		list, err := wishlist.GetWishlist(chatID, username)
		if err != nil {
			msg.Text = beautifulError(err)
		} else {
			wlk := makeWishlistKeyboard(CommandFulfill, username, true, list)
			if len(wlk.InlineKeyboard) == 0 { // when all wishes are fulfilled
				msg.Text = t.G("All your wishes were already fulfilled")
			} else {
				msg.Text = t.G("Which wish of yours do you want to mark as fulfilled?")
				msg.Text += t.G("_(unlisted wishes were already fulfilled)_")
				msg.ReplyMarkup = wlk
				// see handleCallbackQuery to continue with this conversation
			}
		}
	}

	bot.Send(msg)

}
