package wishlist

import (
	"errors"
	"log"

	t "git.leon.wtf/leon/group-wishlist-telegram-bot/translator"
)

func GetWishlist(chatID int64, userInfo UserInfo) (*Wishlist, error) {
	db, err := loadChatDBFile(chatID)
	if err != nil {
		return nil, err
	}
	wishlist, ok := db.Wishlists[userInfo.ID]
	if !ok {
		return nil, &NoWishesForUserError{
			GenericWishlistError{
				Msg: t.G("User %s has not expressed any wishes yet", userInfo.Username),
				Err: nil,
			},
		}
	}
	return wishlist, nil
}

func AddWish(chatID int64, userInfo UserInfo, wish *Wish) error {

	db, err := loadChatDBFile(chatID)
	if err != nil {
		var e *NoDatabaseForChatError
		if errors.As(err, &e) { // create a new DB file if it does not exist already
			db = &chatDBFile{
				ChatID:    chatID,
				Wishlists: make(map[int64]*Wishlist),
			}
		} else {
			return err
		}
	}

	// get a copy of the existing wishlist
	copyWishlist, ok := db.Wishlists[userInfo.ID]
	if !ok { // there is no wishlist for this user
		copyWishlist = &Wishlist{
			Username: userInfo.Username,
			// Wishes will be initializes in the next step
		}
	}
	if copyWishlist.Wishes == nil { // create wishlist if it does not exist already
		copyWishlist.Wishes = make([]*Wish, 0)
	}
	copyWishlist.Wishes = append(copyWishlist.Wishes, wish)
	// check if the username is still the same and update if, if it is not
	if userInfo.Username != copyWishlist.Username {
		log.Printf("Discovered username update: '%s' > '%s'", copyWishlist.Username, userInfo.Username)
		copyWishlist.Username = userInfo.Username
	}
	// assign the new copyWishlist back to the db struct
	db.Wishlists[userInfo.ID] = copyWishlist

	// save the db struct
	if err := db.Save(); err != nil {
		return err
	}
	return nil

}

func FulfillWish(chatID int64, userInfo UserInfo, wishID int) error {
	if wishID < 1 {
		return &WishIDInvalidError{
			GenericWishlistError{
				Msg: t.G("Wish ID cannot be below 1"),
				Err: nil,
			},
		}
	}
	realWishID := wishID - 1 // indexing at 0

	db, err := loadChatDBFile(chatID)
	if err != nil {
		return err
	}
	wishlist, ok := db.Wishlists[userInfo.ID]
	if !ok {
		return &NoWishesForUserError{
			GenericWishlistError{
				Msg: t.G("Wishlist does not exist for user %s", userInfo.Username),
			},
		}
	}
	if len(wishlist.Wishes) <= realWishID {
		return &WishDoesNotExistError{
			GenericWishlistError{
				Msg: t.G("Wish %d does not exist", wishID),
			},
		}
	}
	if wishlist.Wishes[realWishID].Fulfilled {
		return &WishAlreadyFulfilledError{
			GenericWishlistError{
				Msg: t.G("Wish %d is already fulfilled", wishID),
			},
		}
	}

	db.Wishlists[userInfo.ID].Wishes[realWishID].Fulfilled = true
	if err := db.Save(); err != nil {
		return err
	}
	return nil
}

func GetUsersWithWishes(chatID int64) ([]*UserInfo, error) {
	db, err := loadChatDBFile(chatID)
	if err != nil {
		return nil, err
	}
	userInfos := make([]*UserInfo, 0)
	for userID, data := range db.Wishlists {
		userInfos = append(userInfos, &UserInfo{
			ID:       userID,
			Username: data.Username,
		})
	}
	return userInfos, nil
}
