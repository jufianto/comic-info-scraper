package services

type Client struct {
	URLsite  string
	Headfull bool
}

type OptionClient func(*Client)

func NewClient(urlSite string) *Client {
	return &Client{URLsite: urlSite}
}

func WithHeadfull(hf bool) OptionClient {
	return func(c *Client) {
		c.Headfull = hf
	}
}
