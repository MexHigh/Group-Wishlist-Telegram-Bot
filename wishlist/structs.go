package wishlist

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
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
	p := path.Join("db", strconv.FormatInt(db.ChatID, 10)+".json")
	jsonBytes, err := json.MarshalIndent(*db, "", "    ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(p, jsonBytes, 0644); err != nil {
		return err
	}
	return nil
}

func loadChatDBFile(chatID int64) (*chatDBFile, error) {
	p := path.Join("db", strconv.FormatInt(chatID, 10)+".json")
	jsonFile, err := os.Open(p)
	if err != nil {
		return nil, errors.New("chat database does not exist for this chat")
	}
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	db := chatDBFile{}
	if err := json.Unmarshal(jsonBytes, &db); err != nil {
		return nil, err
	}
	return &db, nil
}
