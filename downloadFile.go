package bakeryclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) DownloadFileAsBytes(piId, filename string) ([]byte, error) {
	fullUrl := fmt.Sprintf("%v/oven/%v/download/%v", c.url, piId, filename)
	res, err := http.Get(fullUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("DownloadFileAsBytes returned status code %v", res.StatusCode)
	}

	return ioutil.ReadAll(res.Body)
}
