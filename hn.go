package hn

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
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

func (c *Client) GetItems(ids []int, batchSize int) ([]*Item, error) {
	var wg sync.WaitGroup
	itemCh := make(chan *Item, batchSize)
	errCh := make(chan error, batchSize)
	sem := make(chan struct{}, batchSize)

	for _, id := range ids {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- struct{}{}
			item, err := c.GetItem(id)
			if err != nil {
				errCh <- err
				<-sem
				return
			}
			itemCh <- item
			<-sem
		}(id)
	}

	go func() {
		wg.Wait()
		close(itemCh)
		close(errCh)
	}()

	var items []*Item
	var err error

	for {
		select {
		case item, ok := <-itemCh:
			if !ok {
				itemCh = nil
			} else {
				items = append(items, item)
			}
		case e, ok := <-errCh:
			if !ok {
				errCh = nil
			} else if err == nil {
				err = e
			}
		}

		if itemCh == nil && errCh == nil {
			break
		}
	}

	return items, err
}

type User struct {
	About     string
	Created   time.Time
	ID        string
	Karma     int
	Submitted []int
}

type userResponse struct {
	About     string `json:"about"`
	Created   int64  `json:"created"`
	ID        string `json:"id"`
	Karma     int    `json:"karma"`
	Submitted []int  `json:"submitted"`
}

func (c *Client) GetUser(username string) (*User, error) {
	url := fmt.Sprintf("%s/user/%s.json", c.baseURL, username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status: %d %s", resp.StatusCode, resp.Status)
	}

	var userResp userResponse
	if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	user := &User{
		About:     userResp.About,
		Created:   time.Unix(userResp.Created, 0),
		ID:        userResp.ID,
		Karma:     userResp.Karma,
		Submitted: userResp.Submitted,
	}

	return user, nil
}

func (c *Client) GetTopStoryIDs() ([]int, error) {
	return c.getStoryIDs("topstories")
}

func (c *Client) GetNewStoryIDs() ([]int, error) {
	return c.getStoryIDs("newstories")
}

func (c *Client) GetBestStoryIDs() ([]int, error) {
	return c.getStoryIDs("beststories")
}

func (c *Client) GetAskStoryIDs() ([]int, error) {
	return c.getStoryIDs("askstories")
}

func (c *Client) GetShowStoryIDs() ([]int, error) {
	return c.getStoryIDs("showstories")
}

func (c *Client) GetJobStoryIDs() ([]int, error) {
	return c.getStoryIDs("jobstories")
}

func (c *Client) getStoryIDs(storyType string) ([]int, error) {
	url := fmt.Sprintf("%s/%s.json", c.baseURL, storyType)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status: %d %s", resp.StatusCode, resp.Status)
	}

	var itemIDs []int
	if err := json.NewDecoder(resp.Body).Decode(&itemIDs); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: %v", err)
	}

	return itemIDs, nil
}
