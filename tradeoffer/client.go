package tradeoffer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Philipp15b/go-steam/netutil"
	"net/http"
	"strconv"
)

type APIKey string

const apiUrl = "http://api.steampowered.com/IEconService/%s/v%d"

type Client struct {
	client *http.Client
	key    APIKey
}

func NewClient(key APIKey) *Client {
	return &Client{
		new(http.Client),
		key,
	}
}

func (c *Client) GetOffers() (*TradeOffers, error) {
	resp, err := c.client.Get(fmt.Sprintf(apiUrl, "GetTradeOffers", 1) + "?" + netutil.ToUrlValues(map[string]string{
		"key":                 string(c.key),
		"get_sent_offers":     "1",
		"get_received_offers": "1",
	}).Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	t := new(struct {
		Response *TradeOffers
	})
	err = json.NewDecoder(resp.Body).Decode(t)
	if err != nil {
		return nil, err
	}
	return t.Response, nil
}

type actionResult struct {
	Success bool
	Error   string
}

func (c *Client) action(method string, version uint, id TradeOfferId) error {
	resp, err := c.client.Do(netutil.NewPostForm(fmt.Sprintf(apiUrl, method, version), netutil.ToUrlValues(map[string]string{
		"key":          string(c.key),
		"tradeofferid": strconv.FormatUint(uint64(id), 10),
	})))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New(method + " error: status code not 200")
	}
	return nil
}

func (c *Client) Decline(id TradeOfferId) error {
	return c.action("DeclineTradeOffer", 1, id)
}

func (c *Client) Cancel(id TradeOfferId) error {
	return c.action("CancelTradeOffer", 1, id)
}
