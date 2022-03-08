package translator

var translations map[string]map[Language]string

func init() {
	translations = map[string]map[Language]string{
		// GENERICS START //
		"test": {
			"en": "Test string in english",
			"de": "Test string auf Deutsch",
		},
		"not implemented": {
			"en": "Not implemented yet",
			"de": "Noch nicht implementiert",
		},
		// GENERICS END //

		// BOT COMMANDS START //
		"command_help": {
			"en": "help",
			"de": "hilfe",
		},
		"command_help_desc": {
			"en": "Shows a help message",
			"de": "Zeigt eine Hilfeseite",
		},
		"command_wish": {
			"en": "wish",
			"de": "wunsch",
		},
		"command_wish_desc": {
			"en": "Adds a new wish",
			"de": "Fügt einen neuen Wunsch hinzu",
		},
		"command_wishlist": {
			"en": "wishlist",
			"de": "wunschliste",
		},
		"command_wishlist_desc": {
			"en": "Shows wishes of someone",
			"de": "Zeigt die Wünsche von jemandem an",
		},
		"command_fulfill": {
			"en": "fulfill",
			"de": "verwirklichen",
		},
		"command_fulfill_desc": {
			"en": "Fulfills one own wish",
			"de": "Erfüllt einen eigenen Wunsch",
		},
		// BOT COMMANDS END //

		// NORMAL TRANSLATIONS START //
		"This bot can only be used in group chats": {
			"en": "This bot can only be used in group chats",
			"de": "Dieser Bot kann nur in Gruppenchats verwendet werden",
		},
		"*Wishlist for @%s*": {
			"en": "*Wishlist for @%s*",
			"de": "*Wunschliste für @%s*",
		},
		"Wish %d marked as fulfilled": {
			"en": "Wish %d marked as fulfilled",
			"de": "Wunsch %d als erfüllt markiert",
		},
		"Please provide your wish with your command!": {
			"en": "Please provide your wish with your command!",
			"de": "Bitte gib den Wunsch mit dem Command zusammen an!",
		},
		"Example: `/wish Diamond necklace`": {
			"en": "Example: `/wish Diamond necklace`",
			"de": "Beispiel: `/wunsch Halskette`",
		},
		"Wish created": {
			"en": "Wish created",
			"de": "Wunsch erstellt",
		},
		"Which wishlist do you want to take a look at?": {
			"en": "Which wishlist do you want to take a look at?",
			"de": "Wessen Wunschliste willst du ansehen?",
		},
		"_(users that are not listed do not have any wishes)_": {
			"en": "_(users that are not listed do not have any wishes)_",
			"de": "_(Benutzer, die nicht gelistet sind, haben noch keine Wünsche geäußert)_",
		},
		"All your wishes were already fulfilled": {
			"en": "All your wishes were already fulfilled",
			"de": "Alle deine Wünsche wurden bereits erfüllt",
		},
		"Which wish of yours do you want to mark as fulfilled?": {
			"en": "Which wish of yours do you want to mark as fulfilled?",
			"de": "Welchen deiner Wünsche möchtest du als erfüllt markieren?",
		},
		"_(unlisted wishes were already fulfilled)_": {
			"en": "_(unlisted wishes were already fulfilled)_",
			"de": "_(ungelistete Wünsche wurden bereits erfüllt)_",
		},
		"User @%s has not expressed any wishes yet": {
			"en": "User @%s has not expressed any wishes yet",
			"de": "Benutzer @%s hat bisher noch keine Wünsche geäußert",
		},
		"Wish ID cannot be below 1": {
			"en": "Wish ID cannot be below 1",
			"de": "Wunsch ID kann nicht kleiner als 1 sein",
		},
		"Wishlist does not exist for user @%s": {
			"en": "Wishlist does not exist for user @%s",
			"de": "Wunschlist existiert nicht für Benutzer @%s",
		},
		"Wish %d does not exist": {
			"en": "Wish %d does not exist",
			"de": "Wunsch %d existiert nicht",
		},
		"Wish %d is already fulfilled": {
			"en": "Wish %d is already fulfilled",
			"de": "Wunsch %d wurde bereits erfüllt",
		},
		"No one in this chat has made a wish yet.\nUse `/wish` to add one.": {
			"en": "No one in this chat has made a wish yet.\nUse `/wish` to add one.",
			"de": "Niemand in diesem Chat hat bisher einen Wunsch geäußert.\nBenutze `/wunsch` um einen hinzuzufügen.",
		},
		"_(fulfilled)_": {
			"en": "_(fulfilled)_",
			"de": "_(erfüllt)_",
		},
		// NORMAL TRANSLATIONS END //
	}
}
