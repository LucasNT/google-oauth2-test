package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	createoauthconfig "github.com/LucasNT/google-oauth2-test/internal/createOauthConfig"
)

func main() {
	jsonFilePtr := flag.String("json", "./client_secret.json", "path of the json file that yout can download form gcloud")

	flag.Parse()

	oauth2Config, err := createoauthconfig.New(*jsonFilePtr)
	if err != nil {
		panic(err)
	}

	f, err := os.Open("./token.txt")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	authUrl, err := url.Parse("https://oauth2.googleapis.com/token")
	if err != nil {
		panic(err)
	}

	parameters := url.Values{}
	parameters.Add("client_id", oauth2Config.ClientID)
	parameters.Add("client_secret", oauth2Config.ClientSecret)
	parameters.Add("grant_type", "refresh_token")
	parameters.Add("refresh_token", string(data))
	authUrl.RawQuery = parameters.Encode()

	resp, err := http.Post(authUrl.String(), "", nil)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	a, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(a))

}
