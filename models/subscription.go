package models

import (
	"time"
)

// Subscription defines a feed specific to a user.
type Subscription struct {
	ID     string
	UserID string
	FeedID string
	Title  string // the feed title may be overwritten by the subscriber
	Added  time.Time
	Muted  bool
	Tags   []string
	Notes  string
}

// AddTag adds the tag  to the subscription
func (z *Subscription) AddTag(name string) {
	if !z.HasTag(name) {
		z.Tags = append(z.Tags, name)
	}
}

// HasTag tests if the tag belongs to the subscription
func (z *Subscription) HasTag(name string) bool {
	for _, value := range z.Tags {
		if value == name {
			return true
		}
	}
	return false
}

// RemoveTag removes the tag from the subscription
func (z *Subscription) RemoveTag(name string) {
	for i, value := range z.Tags {
		if value == name {
			z.Tags = append(z.Tags[:i], z.Tags[i+1:]...)
		}
	}
}

// Subscriptions is a collection of Subscription objects
type Subscriptions []*Subscription
