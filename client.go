package canvaslms

import (
	"github.com/go-resty/resty/v2"
	"github.com/tomnomnom/linkheader"
)

// Client -
type Client struct {
	RESTClient *resty.Client
}

// NewClient -
func NewClient(baseURL, token *string) (*Client, error) {
	c := Client{}

	c.RESTClient = resty.New().
		SetBaseURL(*baseURL).
		SetAuthToken(*token).
		SetAuthScheme("Bearer")

	return &c, nil
}

func (c *Client) getResults(URL string) ([]byte, error) {

	resp, err := c.RESTClient.R().Get(URL)
	if err != nil {
		return nil, err
	}

	results := resp.Body()

	links := linkheader.Parse(resp.Header().Get("Link")).FilterByRel("next")
	if len(links) != 0 {
		recurse, err := c.getResults(links[0].URL)
		if err != nil {
			return nil, err
		}
		results = append(results[:len(results)-1], ',')
		results = append(results, recurse[1:]...)
	}
	return results, nil
}
