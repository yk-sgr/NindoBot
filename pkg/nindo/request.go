package nindo

import (
	"io/ioutil"
	"net/url"
)

func (c *Client) get(u *url.URL) ([]byte, error) {
	res, err := c.httpClient.Get(c.baseURL.ResolveReference(u).String())
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}
