package utils

import (
	"net/http"
	"os"
)

func MakeRequest(addedURL string) (*http.Response, error) {
	fullURL := os.Getenv("IMBD_URL") + addedURL
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("IMBD_API_KEY"))

	return http.DefaultClient.Do(req)
}
