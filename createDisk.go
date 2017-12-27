package bakeryclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) CreateDisk(sizeinMb int) (string, error) {
	var diskParams struct {
		Size int
	}

	diskParams.Size = sizeinMb
	jsonBytes, _ := json.Marshal(diskParams)

	req, _ := http.NewRequest("POST", c.url+"/disks", bytes.NewBuffer(jsonBytes))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("CreateDisk returned status code %v", resp.StatusCode)
	}

	var disk Disk
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &disk)
	if err != nil {
		return "", err
	}

	return disk.ID, nil
}
