package nindo

import (
	"net/url"
	"time"
)

type ViralPostType string

var (
	ViralPostTypeLikes         ViralPostType = "likes"
	ViralPostTypeComments      ViralPostType = "kommentare"
	ViralPostTypeViews         ViralPostType = "views"
	ViralPostTypeRetweets      ViralPostType = "retweets"
	ViralPostTypeMaxViewers    ViralPostType = "max. zuschauer"
	ViralPostTypeLongestStream ViralPostType = "längster stream"
)

type ViralPost struct {
	PostID    string        `json:"postID"`
	Platform  Platform      `json:"platform"`
	Type      string        `json:"type"`
	Value     int           `json:"value"`
	Timestamp time.Time     `json:"timestamp"`
	Data      ViralPostData `json:"_post"`
}

type ViralPostData struct {
	Title          string           `json:"title"`
	ContentChecked bool             `json:"contentChecked"`
	FSK18          bool             `json:"FSK18"`
	Clickbait      string           `json:"clickbait"`
	Shitstorm      string           `json:"shitstorm"`
	Ad             string           `json:"ad"`
	Channel        ViralPostChannel `json:"_channel"`
}

type ViralPostChannel struct {
	ChannelID    string                 `json:"channelID"`
	Avatar       string                 `json:"avatar"`
	CachedAvatar bool                   `json:"cachedAvatar"`
	LastPostID   string                 `json:"lastPostID"`
	Artist       ViralPostChannelArtist `json:"_artist"`
}

type ViralPostChannelArtist struct {
	Name string `json:"name"`
	ID   string `json:"_id"`
}

func jsonParseViralPosts(b []byte) ([]*ViralPost, error) {
	var posts []*ViralPost
	err := json.Unmarshal(b, &posts)
	return posts, err
}

func (c *Client) GetViralPosts() ([]*ViralPost, error) {
	b, err := c.get(&url.URL{Path: "/viral"})
	if err != nil {
		return nil, err
	}
	return jsonParseViralPosts(b)
}
