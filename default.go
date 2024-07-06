package hn

var _defaultClient = New()

func GetItem(id int) (*Item, error) {
	return _defaultClient.GetItem(id)
}

func GetItems(ids []int) ([]*Item, error) {
	return _defaultClient.GetItems(ids)
}

func GetUser(username string) (*User, error) {
	return _defaultClient.GetUser(username)
}

// GetTopStoryIDs fetches the IDs of the top stories from Hacker News.
func GetTopStoryIDs() ([]int, error) {
	return _defaultClient.GetTopStoryIDs()
}

// GetNewStoryIDs fetches the IDs of the new stories from Hacker News.
func GetNewStoryIDs() ([]int, error) {
	return _defaultClient.GetNewStoryIDs()
}

// GetBestStoryIDs fetches the IDs of the best stories from Hacker News.
func GetBestStoryIDs() ([]int, error) {
	return _defaultClient.GetBestStoryIDs()
}

// GetAskStoryIDs fetches the IDs of the ask stories from Hacker News.
func GetAskStoryIDs() ([]int, error) {
	return _defaultClient.GetAskStoryIDs()
}

// GetShowStoryIDs fetches the IDs of the show stories from Hacker News.
func GetShowStoryIDs() ([]int, error) {
	return _defaultClient.GetShowStoryIDs()
}

// GetJobStoryIDs fetches the IDs of the job stories from Hacker News.
func GetJobStoryIDs() ([]int, error) {
	return _defaultClient.GetJobStoryIDs()
}

// WithConcurrency returns a new Client with the specified concurrency level.
func WithConcurrency(n int) *Client {
	return _defaultClient.WithConcurrency(n)
}
