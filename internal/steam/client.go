package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	apiKey      string
	httpsClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:      apiKey,
		httpsClient: &http.Client{},
	}
}

type GameResult struct {
	AppID int    `json:"appid"`
	Name  string `json:"name"`
}

func (c *Client) SearchGames(steamID string) ([]GameResult, error) {
	// url.QueryEscape is used to ensure the steamID is properly encoded for the URL

	endpoint := fmt.Sprintf(
		"https://store.steampowered.com/api/storesearch/?term=%s&cc=us&l=en",
		url.QueryEscape(steamID),
	)
	resp, err := c.httpsClient.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("steam API request failed: %w", err)
	}
	defer resp.Body.Close()

	var result struct {
		Total int `json:"total"`
		Items []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("steam search decode failed: %w", err)
	}

	games := make([]GameResult, 0, len(result.Items))
	for _, item := range result.Items {
		games = append(games, GameResult{
			AppID: item.ID,
			Name:  item.Name,
		})
	}

	return games, nil
}
