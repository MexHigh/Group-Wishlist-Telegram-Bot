# Group Wishlist Telegram bot

This Telegram bot keeps track of wishes and wishlists inside a group chat. 
Every Group in which the bot is a member has a database file, where user wishes are tracked.
Users inside the group can query the wishlists of all other users in the group that placed wishes.
This makes this bot **perfect for family and friends groups**!

The bot cannot be used in private chats or supergroups.

## Commands

| Command (in English) | Description                                                                |
| -------------------- | -------------------------------------------------------------------------- |
| `/help`              | Shows a help page                                                          |
| `/wish <wish>`       | Place a wish                                                               |
| `/wishlist`          | Displays the wishlist of another user                                      |
| `/fulfill`           | Fulfills an own wish (Wishes of other users cannot be marked as fulfilled) |


## Translation

The bots messages are completely translatable (including the commands!). Just add your translations to `translator/translations.go`. 
Currently, the languages English (`'en'`) and German (`'de'`) are available. You can set the bots language with `--language 'de'|'en''`.


## Compile and run

1. Clone this repo: `git clone https://git.leon.wtf/leon/group-wishlist-telegram-bot`
2. Compile with Go (at least `v1.17`): `go build .`
3. Copy the `config.example.json` to `config.json` and paste in your Bot token
4. Run the bot with `./group-wishlist-telegram-bot --config ./config.json --language en`


## Run with Docker

1. Grab the `docker-compose.yml` file and adjust it to your likings
2. Copy the `config.example.json` to `config.json` and paste in your Bot token
3. Create the folder named `db/`
4. Run `docker-compose up -d` (uses the config from step 2 and persists wishes in `db/`)
5. See logs with `docker-compose logs -f`