package wishlist

import (
	"errors"
	"fmt"
)

func GetWishlist(chatID int64, username Username) (Wishlist, error) {
	db, err := loadChatDBFile(chatID)
	if err != nil {
		return nil, err
	}
	wishlist, ok := db.Wishes[username]
	if !ok {
		return nil, errors.New("wishlist does not exist for user " + string(username))
	}
	return wishlist, nil
}

func GetWish(chatID int64, username Username, wishID int) *Wish {
	return nil // TODO not implemented (unused)
}

func AddWish(chatID int64, username Username, wish *Wish) error {
	db, err := loadChatDBFile(chatID)
	if err != nil {
		// create a new DB file if it does not exist already
		db = &chatDBFile{
			ChatID: chatID,
			Wishes: make(map[Username]Wishlist),
		}
	}
	db.Wishes[username] = append(db.Wishes[username], wish)
	db.Save()
	return nil
}

func FulfillWish(chatID int64, username Username, wishID int) error {

	if wishID < 1 {
		return errors.New("wish ID cannot be below 1")
	}
	realWishID := wishID - 1 // indexing at 0

	db, err := loadChatDBFile(chatID)
	if err != nil {
		return err
	}
	wishes, ok := db.Wishes[username]
	if !ok {
		return fmt.Errorf("wishlist does not exist for user %s", string(username))
	}
	if len(wishes) <= realWishID {
		return fmt.Errorf("wish %d does not exist", wishID)
	}
	if wishes[realWishID].Fulfilled {
		return fmt.Errorf("wish %d is already fulfilled", wishID)
	}

	db.Wishes[username][realWishID].Fulfilled = true
	db.Save()
	return nil

}
