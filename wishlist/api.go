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
		return nil, &NoWishesForUserError{
			GenericWishlistError{
				Msg: fmt.Sprintf("User %s has not expressed any wishes yet", string(username)),
				Err: nil,
			},
		}
	}
	return wishlist, nil
}

func AddWish(chatID int64, username Username, wish *Wish) error {
	db, err := loadChatDBFile(chatID)
	if err != nil {
		var e *NoDatabaseForChatError
		if errors.As(err, &e) { // create a new DB file if it does not exist already
			db = &chatDBFile{
				ChatID: chatID,
				Wishes: make(map[Username]Wishlist),
			}
		} else {
			return err
		}
	}
	db.Wishes[username] = append(db.Wishes[username], wish)
	if err := db.Save(); err != nil {
		return err
	}
	return nil
}

func FulfillWish(chatID int64, username Username, wishID int) error {
	if wishID < 1 {
		return &WishIDInvalidError{
			GenericWishlistError{
				Msg: "Wish ID cannot be below 1",
				Err: nil,
			},
		}
	}
	realWishID := wishID - 1 // indexing at 0

	db, err := loadChatDBFile(chatID)
	if err != nil {
		return err
	}
	wishes, ok := db.Wishes[username]
	if !ok {
		return &NoWishesForUserError{
			GenericWishlistError{
				Msg: fmt.Sprintf("Wishlist does not exist for user %s", string(username)),
			},
		}
	}
	if len(wishes) <= realWishID {
		return &WishDoesNotExistError{
			GenericWishlistError{
				Msg: fmt.Sprintf("Wish %d does not exist", wishID),
			},
		}
	}
	if wishes[realWishID].Fulfilled {
		return &WishAlreadyFulfilledError{
			GenericWishlistError{
				Msg: fmt.Sprintf("Wish %d is already fulfilled", wishID),
			},
		}
	}

	db.Wishes[username][realWishID].Fulfilled = true
	if err := db.Save(); err != nil {
		return err
	}
	return nil
}

func GetUsersWithWishes(chatID int64) ([]Username, error) {
	db, err := loadChatDBFile(chatID)
	if err != nil {
		return nil, err
	}
	users := make([]Username, 0)
	for user := range db.Wishes {
		users = append(users, user)
	}
	return users, nil
}
