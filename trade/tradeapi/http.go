package tradeapi

import (
	"encoding/json"
	"github.com/Philipp15b/go-steam/netutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func (t *Trade) setCookies(sessionId, steamLogin string) {
	t.client.Jar, _ = cookiejar.New(new(cookiejar.Options))
	base, err := url.Parse(cookiePath)
	if err != nil {
		panic(err)
	}
	t.client.Jar.SetCookies(base, []*http.Cookie{
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

// Ajax POSTs to an API endpoint that should return a status
func (t *Trade) postWithStatus(url string, data map[string]string) (*Status, error) {
	status := new(Status)

	req := netutil.NewPostForm(url, netutil.ToUrlValues(data))
	// Tales of Madness and Pain, Episode 1: If you forget this, Steam will return an error
	// saying "missing required parameter", even though they are all there. IT WAS JUST THE HEADER, ARGH!
	req.Header.Add("Referer", t.baseUrl)

	resp, err := t.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(status)
	if err != nil {
		return nil, err
	}
	return status, nil
}
