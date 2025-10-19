package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/LucasNT/google-oauth2-test/webServer"
	"github.com/skratchdot/open-golang/open"
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

func main() {

	jsonFilePtr := flag.String("json", "./client_secret.json", "path of the json file that yout can download form gcloud")

	flag.Parse()

	authConfig, err := parserJson(*jsonFilePtr)
	if err != nil {
		panic(err)
	}

	var oauth2Config *oauth2.Config = &oauth2.Config{
		ClientID:     authConfig["web"].ClientID,
		ClientSecret: authConfig["web"].ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:3000/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}

	webServer, err := webServer.New(oauth2Config)
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr:    ":3000",
		Handler: webServer,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("Making oauth request")

	open.Run("http://localhost:3000/")

	time.Sleep(400 * time.Second)
}
