package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	createoauthconfig "github.com/LucasNT/google-oauth2-test/internal/createOauthConfig"
	requestuserinfo "github.com/LucasNT/google-oauth2-test/internal/requestUserInfo"
	"github.com/LucasNT/google-oauth2-test/internal/webServer"
	"github.com/skratchdot/open-golang/open"
)

func main() {

	jsonFilePtr := flag.String("json", "./client_secret.json", "path of the json file that yout can download form gcloud")

	flag.Parse()

	oauth2Config, err := createoauthconfig.New(*jsonFilePtr)

	webServer, ch, err := webServer.New(oauth2Config)
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

	token := <-ch

	out, err := requestuserinfo.RequestUserInfo(token)

	fmt.Println(string(out))
}
