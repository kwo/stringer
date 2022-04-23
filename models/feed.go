package models

import (
	"time"
)

// Feed feed descriptor
type Feed struct {
	ID           string
	URL          string
	SiteURL      string
	IconURL      string
	ETag         string
	LastModified time.Time
	LastUpdated  time.Time
}

// Feeds is a collection of Feed elements
type Feeds []*Feed

func (z Feeds) GetByID(id string) *Feed {
	for _, feed := range z {
		if feed.ID == id {
			return feed
		}
	}
	return nil
}
