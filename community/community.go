package community

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const cookiePath = "http://steamcommunity.com/"

func SetCookies(client *http.Client, sessionId, steamLogin string) {
	if client.Jar == nil {
		client.Jar, _ = cookiejar.New(new(cookiejar.Options))
	}
	base, err := url.Parse(cookiePath)
	if err != nil {
		panic(err)
	}
	client.Jar.SetCookies(base, []*http.Cookie{
		&http.Cookie{
			Name:  "sessionid",
			Value: sessionId,
		},
		&http.Cookie{
			Name:  "steamLogin",
			Value: steamLogin,
		},
	})
}
