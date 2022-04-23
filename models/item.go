package models

import (
	"time"
)

// Item from a feed
type Item struct {
	ID       string
	FeedID   string
	Created  time.Time
	Updated  time.Time
	URL      string
	Author   string
	Title    string
	Summary  string
	Contents string
}

type Items []*Item

func (z Items) GetByID(id string) *Item {
	for _, item := range z {
		if item.ID == id {
			return item
		}
	}
	return nil
}
