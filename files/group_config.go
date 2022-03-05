package files

// TEMP

import "time"

type GroupMember string

type Wishlist struct {
	WishedAt *time.Time `json:"wished_at"`
	Wish     string     `json:"wish"`
}

type GroupConfig struct {
	GroupID string                    `json:"group_id"`
	Members []*GroupMember            `json:"members"`
	Wishes  map[GroupMember]*Wishlist `json:"wishes"`
}
