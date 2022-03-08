package wishlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type Username string

type Wish struct {
	WishedAt  time.Time `json:"wished_at"`
	Wish      string    `json:"wish"`
	Fulfilled bool      `json:"fulfilled"`
}

type Wishlist []*Wish

func (w *Wishlist) String() (s string) {
	for i, wish := range *w {
		s += fmt.Sprintf("*%d.* %s", i+1, wish.Wish)
		if wish.Fulfilled {
			s += " _(fulfilled)_"
		}
		s += "\n"
	}
	return
}

type chatDBFile struct {
	ChatID int64                 `json:"group_id"`
	Wishes map[Username]Wishlist `json:"wishes"`
}

func (db *chatDBFile) Save() error {
	f := strconv.FormatInt(db.ChatID, 10) + ".json"
	f = strings.ReplaceAll(f, "-", "m") // replace negative sign with 'm'
	p := path.Join("db", f)
	jsonBytes, err := json.MarshalIndent(*db, "", "    ")
	if err != nil {
		return &InternalError{
			GenericWishlistError{
				Err: err,
			},
		}
	}
	if err := os.WriteFile(p, jsonBytes, 0644); err != nil {
		return &InternalError{
			GenericWishlistError{
				Err: err,
			},
		}
	}
	return nil
}

func loadChatDBFile(chatID int64) (*chatDBFile, error) {
	f := strconv.FormatInt(chatID, 10) + ".json"
	f = strings.ReplaceAll(f, "-", "m") // replace negative sign with 'm'
	p := path.Join("db", f)
	jsonFile, err := os.Open(p)
	if err != nil {
		return nil, &NoDatabaseForChatError{
			GenericWishlistError{
				Msg: "No one in this chat has made a wish yet.\nUse `/wish` to add one.",
				Err: err,
			},
		}
	}
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, &InternalError{
			GenericWishlistError{
				Err: err,
			},
		}
	}
	db := chatDBFile{}
	if err := json.Unmarshal(jsonBytes, &db); err != nil {
		return nil, &InternalError{
			GenericWishlistError{
				Err: err,
			},
		}
	}
	return &db, nil
}
