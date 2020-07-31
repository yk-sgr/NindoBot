package nindo

import (
	"net/url"
)

type RankArtist struct {
	ID       string `json:"id"`
	ArtistID string `json:"artistID"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	UserID   string `json:"userID"`
	// _artist ???
	Rank        int    `json:"rank"`
	Value       int    `json:"value"`
	ArtistName  string `json:"artistName"`
	ShowChannel bool   `json:"showChannel"`
}

func jsonParseRankArtists(b []byte) ([]*RankArtist, error) {
	var artists []*RankArtist
	err := json.Unmarshal(b, &artists)
	return artists, err
}

func (c *Client) GetTikTokRankChartsByLikes() ([]*RankArtist, error) {
	return c.getRanking(&url.URL{Path: "/ranks/charts/tiktok/rankLikes/small"})
}

func (c *Client) GetTwitchRankChartsByViewers() ([]*RankArtist, error) {
	return c.getRanking(&url.URL{Path: "/ranks/charts/twitch/rankViewer/small"})
}

func (c *Client) GetTwitterRankChartsByLikes() ([]*RankArtist, error) {
	return c.getRanking(&url.URL{Path: "/ranks/charts/twitter/rankLikes/small"})
}

func (c *Client) GetInstagramRankChartsByLikes() ([]*RankArtist, error) {
	return c.getRanking(&url.URL{Path: "/ranks/charts/instagram/rankLikes/small"})
}

func (c *Client) GetYouTubeRankChartsByViews() ([]*RankArtist, error) {
	return c.getRanking(&url.URL{Path: "/ranks/charts/youtube/rankViews/small"})
}

func (c *Client) getRanking(u *url.URL) ([]*RankArtist, error) {
	b, err := c.get(u)
	if err != nil {
		return nil, err
	}
	return jsonParseRankArtists(b)
}
