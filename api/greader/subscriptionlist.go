package greader

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kwo/stringer/models"
)

type SubscriptionListProvider interface {
	GetSubscriptionsForUser(string) (models.Subscriptions, error)
	GetTagsForUser(string) (models.Tags, error)
	GetFeeds(...string) (models.Feeds, error)
}

type SubscriptionList struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

type Subscription struct {
	ID         string     `json:"id"`
	Title      string     `json:"title"`
	URL        string     `json:"url"`
	SiteURL    string     `json:"htmlUrl"`
	IconURL    string     `json:"iconUrl"`
	Sort       string     `json:"sortid,omitempty"`
	AddedMS    int64      `json:"firstitemmsec,omitempty"`
	Categories []Category `json:"categories"`
}

type Category struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

func subscriptionList(subscriptionListProvider SubscriptionListProvider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if user := getUserFromContext(ctx); user != nil {

			subscriptions, err := subscriptionListProvider.GetSubscriptionsForUser(user.ID)
			if err != nil {
				log.Printf("subscription list: cannot get subscriptions: %s", err)
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			tags, err := subscriptionListProvider.GetTagsForUser(user.ID)
			if err != nil {
				log.Printf("subscription list: cannot get tags: %s", err)
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			feedIDs := makeFeedIDs(subscriptions)
			feeds, err := subscriptionListProvider.GetFeeds(feedIDs...)
			if err != nil {
				log.Printf("subscription list: cannot get feeds: %s", err)
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			subscriptionList := &SubscriptionList{
				Subscriptions: makeSubscriptions(subscriptions, tags, feeds),
			}

			data, err := json.Marshal(subscriptionList)
			if err != nil {
				log.Printf("subscription list: cannot marshal response: %s", err)
				http.Error(w, "", http.StatusInternalServerError)
				return
			}

			w.Header().Set(hContentType, mimetypeJson)
			_, _ = w.Write(data)
			return

		}
		http.Error(w, "", http.StatusUnauthorized)
	}
}

func makeFeedIDs(subscriptions models.Subscriptions) []string {
	var result []string
	for _, subscription := range subscriptions {
		result = append(result, subscription.FeedID)
	}
	return result
}

func makeSubscriptions(subscriptions models.Subscriptions, tags models.Tags, feeds models.Feeds) []Subscription {

	var result []Subscription

	for _, subscription := range subscriptions {

		feed := feeds.GetByID(subscription.FeedID)
		if feed == nil {
			log.Printf("skipping subscription, no feed: %v", subscription)
			continue
		}

		sub := Subscription{
			ID:      subscription.ID,
			Title:   subscription.Title,
			URL:     feed.URL,
			SiteURL: feed.SiteURL,
			IconURL: feed.IconURL,
			// Sort:       subscription.Title, // TODO: sort
			// AddedMS:    subscription.Added.Unix() * 1000,
			Categories: makeCategories(subscription.Tags, tags),
		}

		result = append(result, sub)

	}

	return result

}

func makeCategories(tagNames []string, tags models.Tags) []Category {
	var result []Category
	for _, tagName := range tagNames {
		tag := tags.GetByName(tagName)
		if tag == nil {
			log.Printf("skipping category, no tag: %s", tagName)
			continue
		}
		result = append(result, Category{
			ID:    fmt.Sprintf("user/-/label/%s", tag.Name),
			Label: tag.Name,
		})
	}
	return result
}
