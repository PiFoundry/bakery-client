package bakeryclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (c *Client) IsPiBaked(piId string) (bool, error) {
	resp, err := c.httpClient.Get(c.url + "/oven/" + piId)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	var parsedResponse PiInfo
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &parsedResponse)
	if err != nil {
		return false, err
	}

	if parsedResponse.Status != INUSE {
		return false, nil
	}

	return true, nil
}
