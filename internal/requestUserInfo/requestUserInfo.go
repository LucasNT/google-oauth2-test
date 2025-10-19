package requestuserinfo

import (
	"fmt"
	"io"
	"net/http"
)

func RequestUserInfo(token string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		return "", fmt.Errorf("Failed to request userinfo %w", err)
	}

	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to request userinfo %w", err)
	}
	defer res.Body.Close()

	out, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Failed to request userinfo %w", err)
	}
	return string(out), nil
}
