package bakeryclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetDisks() (Disks, error) {
	var diskResponse Disks

	resp, err := c.httpClient.Get(c.url + "/disks")
	if err != nil {
		return Disks{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Disks{}, fmt.Errorf("GetDisks returned status code %v", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&diskResponse)
	if err != nil {
		return Disks{}, err
	}

	return diskResponse, nil
}
