package createoauthconfig

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type googleCloudOauthSecret struct {
	ClientID                string   `json:"client_id"`
	ClientSecret            string   `json:"client_secret"`
	ProjectID               string   `json:"project_id"`
	AuthUri                 string   `json:"auth_uri"`
	TokenUri                string   `json:"token_uri"`
	AuthProviderX509CertUrl string   `json:"auth_provider_x509_cert_url"`
	RedirectUris            []string `json:"redirect_uris"`
}

func parserJson(path string) (map[string]googleCloudOauthSecret, error) {
	a := make(map[string]googleCloudOauthSecret)
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file for json secrets parser %w", err)
	}
	defer f.Close()
	fileData, err := io.ReadAll(f)
	err = json.Unmarshal(fileData, &a)
	return a, nil
}

func New(path string) (*oauth2.Config, error) {
	authConfig, err := parserJson(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to parser %s file %w", path, err)
	}

	var oauth2Config *oauth2.Config = &oauth2.Config{
		ClientID:     authConfig["web"].ClientID,
		ClientSecret: authConfig["web"].ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:3000/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
	}

	return oauth2Config, nil
}
