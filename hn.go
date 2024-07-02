package hn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const BaseUrlv0 = "https://hacker-news.firebaseio.com/v0"

type Client struct {
	baseURL string
}

func New() Client {
	return Client{baseURL: BaseUrlv0}
}

type Item struct {
	ID          int
	By          string
	Title       string
	URL         string
	Score       int
	Time        time.Time
	Descendants int
	Kids        []int
	Type        string

	Parent int
	Text   string
	Parts  []int
	Poll   int
}

type itemResponse struct {
	ID          int    `json:"id"`
	By          string `json:"by"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	Score       int    `json:"score"`
	Time        int64  `json:"time"`
	Descendants int    `json:"descendants"`
	Kids        []int  `json:"kids"`
	Type        string `json:"type"`

	Parent int    `json:"parent"`
	Text   string `json:"text"`
	Parts  []int  `json:"parts"` // only in "type": "poll"
	Poll   int    `json:"poll"`  // only in "type": "pollopt"
}

func (c *Client) GetItem(id int) (*Item, error) {
	url := fmt.Sprintf("%s/item/%d.json", c.baseURL, id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status: %d %s", resp.StatusCode, resp.Status)
	}

	var itemResp itemResponse
	if err := json.NewDecoder(resp.Body).Decode(&itemResp); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	item := &Item{
		ID:          itemResp.ID,
		By:          itemResp.By,
		Title:       itemResp.Title,
		URL:         itemResp.URL,
		Score:       itemResp.Score,
		Time:        time.Unix(itemResp.Time, 0),
		Descendants: itemResp.Descendants,
		Kids:        itemResp.Kids,
		Type:        itemResp.Type,
		Parent:      itemResp.Parent,
		Text:        itemResp.Text,
		Parts:       itemResp.Parts,
		Poll:        itemResp.Poll,
	}

	return item, nil
}
