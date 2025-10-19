package webServer

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

type CustomServerHandler struct {
	config  *oauth2.Config
	handler http.Handler
}

func New(config *oauth2.Config) (*CustomServerHandler, error) {
	ret := new(CustomServerHandler)
	ret.config = config

	templ, err := template.New("index.html").Funcs(
		template.FuncMap{
			"join": strings.Join,
		},
	).ParseFiles("./index.html")

	if err != nil {
		return nil, fmt.Errorf("Failed to create template %w", err)
	}

	mut := http.NewServeMux()

	mut.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		authUrl, err := ret.GenerateGoogleAuthURL()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		if err := templ.Execute(w, authUrl); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	})

	mut.HandleFunc("GET /callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		fmt.Println(code)
		token, err := ret.config.Exchange(context.Background(), code)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Println("Access Token", token.AccessToken)
		fmt.Println("Expires In", token.ExpiresIn)
		fmt.Println("Token Type", token.TokenType)
		fmt.Println("Refresh Token", len(token.RefreshToken))
		fmt.Println("Expiry", token.Expiry)
		http.Redirect(w, r, "https://google.com.br", http.StatusSeeOther)
	})

	ret.handler = mut

	return ret, nil
}

func (c *CustomServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.handler.ServeHTTP(w, r)
}

func (c *CustomServerHandler) GenerateGoogleAuthURL() (string, error) {
	authUrl, err := url.Parse("https://accounts.google.com/o/oauth2/v2/auth")
	if err != nil {
		return "", fmt.Errorf("Failed to parse url, %w")
	}
	parameters := url.Values{}
	parameters.Add("scope", strings.Join(c.config.Scopes, " "))
	parameters.Add("client_id", c.config.ClientID)
	parameters.Add("redirect_uri", c.config.RedirectURL)
	parameters.Add("response_type", "code")
	authUrl.RawQuery = parameters.Encode()

	return authUrl.String(), err
}
