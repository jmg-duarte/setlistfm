package setlistfm

type Client struct {
	APIKey string
}

func NewClient(APIKey string) *Client {
	return &Client{APIKey: APIKey}
}
