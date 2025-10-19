# google-oauth2-test

Simple program to test the google oauth2 login.

You will need a google cloud account, to create a oauth2 client. When the client is created will have a button
to download the client secret in json format, that is the file that this program uses to run

The redirect_uri need to be `http://localhost:3000`

# How to run

```bash
go run ./main.go -json=<path to client_secret>
```
