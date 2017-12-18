package bakeryclient

import (
	"fmt"
	"net/http"
)

func (c *Client) RebootPi(piId string) error {
	fullUrl := fmt.Sprintf("%v/oven/%v/reboot", c.url, piId)
	resp, err := c.httpClient.Post(fullUrl, "application/json", nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Reboot Pi returned status code %v", resp.StatusCode)
	}

	return nil
}
