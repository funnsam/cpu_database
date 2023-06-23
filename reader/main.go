package reader

import (
	"fmt"
	"io"
	"net/http"
)

func ReadDatabaseContent() (*string, error) {
	resp, err := http.Get("https://gdirect.cc/d/4tDHR&type=1")
	if err != nil {
		return nil, fmt.Errorf("Unable to get database: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to read database: %v", err)
	}
	b := string(body)
	return &b, nil
}
