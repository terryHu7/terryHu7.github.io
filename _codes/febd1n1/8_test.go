package main

import (
	"fmt"
	"net/url"
	"testing"
)

func TestURL(t *testing.T) {
	// newurl, _ := addQueryParams("http://10.12.224.53:6689/mediafile/getMediaFileList?limit=5&offset=5",
	// 	map[string]string{
	// 		"mode":   "9",
	// 		"limit":  "9",
	// 		"offset": "9",
	// 	})
	url := "http://10.12.224.53:6689/mediafile/getMediaFileList?limit=5&offset=5"
	newport := 6688
	newurl, _ := changePort(url, newport)

	fmt.Println(newurl)
}

func addQueryParams(urlStr string, params map[string]string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	return u.String(), nil
}

func changePort(urlStr string, port int) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	u.Host = fmt.Sprintf("%s:%d", u.Hostname(), port)

	return u.String(), nil
}
