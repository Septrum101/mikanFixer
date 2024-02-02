package main

import (
	"fmt"
	"io"
	"net/http"
)

func fetchRss(token string) ([]byte, error) {
	resp, err := http.Get("https://mikanani.me/RSS/MyBangumi?token=" + token)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch RSS content: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
