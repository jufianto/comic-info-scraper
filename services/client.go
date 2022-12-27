package services

type Client struct {
	URLsite  string
	Headfull bool
}

func NewClient(urlSite string, optionClient ...OptionClient) *Client {
	client := &Client{URLsite: urlSite}
	for _, options := range optionClient {
		options(client)
	}
	return client
}

type OptionClient func(*Client)

func WithHeadfull(hf bool) OptionClient {
	return func(c *Client) {
		c.Headfull = hf
	}
}

func (c *Client) GetHeadfullStatus() bool {
	return c.Headfull
}

type InfoComic struct {
	Title       string
	LastChapter string
}
