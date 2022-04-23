package bogus

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/kwo/stringer/models"
)

func New() *Provider {
	return &Provider{}
}

type Provider struct{}

func (z *Provider) Authenticate(username, password string) (string, bool, error) {
	if username == "hello" && password == "world" {
		token, err := generateToken()
		return token, true, err
	}
	return "", false, nil
}

func (z *Provider) AuthenticateToken(token string) (*models.User, error) {
	u := &models.User{
		ID:           "user/12345",
		Username:     "hello",
		Name:         "KarlO",
		PasswordHash: "243261243130244d524a6b5a377631732e6957747436506d43425864656e495047654d6f68306458736662444b6a4b2f4d64746f583549666e524975",
		Created:      time.Date(2019, time.October, 24, 23, 45, 0, 0, time.UTC),
		Updated:      time.Date(2019, time.October, 24, 23, 45, 0, 0, time.UTC),
	}
	return u, nil
}

func (z *Provider) GetFeeds(feedIDs ...string) (models.Feeds, error) {

	var feeds models.Feeds

	feeds = append(feeds, &models.Feed{
		ID:           "feed/abcde",
		URL:          "https://www.nzz.ch/recent.rss",
		SiteURL:      "https://www.nzz.ch/",
		IconURL:      "https://www.nzz.ch/logo.png",
		ETag:         "etag12345",
		LastModified: time.Date(2019, time.October, 26, 16, 30, 0, 0, time.UTC),
		LastUpdated:  time.Date(2019, time.October, 26, 16, 31, 0, 0, time.UTC),
	})

	feeds = append(feeds, &models.Feed{
		ID:           "feed/fghij",
		URL:          "http://www.spiegel.de/schlagzeilen/eilmeldungen/index.rss",
		SiteURL:      "https://www.spiegel.de/",
		IconURL:      "https://www.spiegel.de/favicon.ico",
		ETag:         "etag67890",
		LastModified: time.Date(2019, time.October, 26, 16, 30, 0, 0, time.UTC),
		LastUpdated:  time.Date(2019, time.October, 26, 16, 31, 0, 0, time.UTC),
	})

	feeds = append(feeds, &models.Feed{
		ID:           "feed/klmnop",
		URL:          "http://feeds.reuters.com/reuters/topNews",
		SiteURL:      "https://www.reuters.com/",
		IconURL:      "https://s3.reutersmedia.net/resources_v2/images/favicon/favicon.ico",
		ETag:         "etag13579",
		LastModified: time.Date(2019, time.October, 26, 16, 30, 0, 0, time.UTC),
		LastUpdated:  time.Date(2019, time.October, 26, 16, 31, 0, 0, time.UTC),
	})

	return feeds, nil

}

func (z *Provider) GetSubscriptionsForUser(userID string) (models.Subscriptions, error) {

	var subscriptons models.Subscriptions

	subscriptons = append(subscriptons, &models.Subscription{
		ID:     "sub/123",
		UserID: "user/12345",
		FeedID: "feed/abcde",
		Title:  "NZZ",
		Added:  time.Date(2019, time.October, 26, 15, 54, 0, 0, time.UTC),
		Muted:  false,
		Tags:   []string{"abc", "xyz"},
		Notes:  "notes nzz",
	})

	subscriptons = append(subscriptons, &models.Subscription{
		ID:     "sub/456",
		UserID: "user/12345",
		FeedID: "feed/fghij",
		Title:  "Spiegel",
		Added:  time.Date(2019, time.October, 26, 15, 54, 0, 0, time.UTC),
		Muted:  false,
		Tags:   []string{"xyz"},
		Notes:  "notes spiegel",
	})

	subscriptons = append(subscriptons, &models.Subscription{
		ID:     "sub/789",
		UserID: "user/12345",
		FeedID: "feed/klmnop",
		Title:  "Reuters",
		Added:  time.Date(2019, time.October, 26, 15, 54, 0, 0, time.UTC),
		Muted:  false,
		Tags:   []string{"abc"},
		Notes:  "notes reuters",
	})

	return subscriptons, nil

}

func (z *Provider) GetTagsForUser(userID string) (models.Tags, error) {

	var tags models.Tags

	tags = append(tags, &models.Tag{
		ID:   "tag/abc",
		Name: "abc",
	})

	tags = append(tags, &models.Tag{
		ID:   "tag/xyz",
		Name: "xyz",
	})

	return tags, nil
}

func (z *Provider) GetItemsUnread(max int) (models.Items, error) {
	var items models.Items
	return items, nil
}

func generateToken() (string, error) {
	c := 16
	b := make([]byte, c)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	token := hex.EncodeToString(b)
	return token, nil
}
