package tradeoffer

import (
	"encoding/json"
	"fmt"
	"github.com/Philipp15b/go-steam/community"
	"github.com/Philipp15b/go-steam/economy/inventory"
	"github.com/Philipp15b/go-steam/netutil"
	"github.com/Philipp15b/go-steam/steamid"
	"net/http"
	"strconv"
	"time"
)

type APIKey string

const apiUrl = "https://api.steampowered.com/IEconService/%s/v%d"

type Client struct {
	client    *http.Client
	key       APIKey
	sessionId string
}

func NewClient(key APIKey, sessionId, steamLogin, steamLoginSecure string) *Client {
	c := &Client{
		new(http.Client),
		key,
		sessionId,
	}
	community.SetCookies(c.client, sessionId, steamLogin, steamLoginSecure)
	return c
}

func (c *Client) GetOffer(offerId uint64) (*TradeOfferResult, error) {
	resp, err := c.client.Get(fmt.Sprintf(apiUrl, "GetTradeOffer", 1) + "?" + netutil.ToUrlValues(map[string]string{
		"key":          string(c.key),
		"tradeofferid": strconv.FormatUint(offerId, 10),
		"language":     "en_us",
	}).Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	t := new(struct {
		Response *TradeOfferResult
	})
	if err = json.NewDecoder(resp.Body).Decode(t); err != nil {
		return nil, err
	}
	return t.Response, nil
}

func (c *Client) GetOffers(getSent bool, getReceived bool, getDescriptions bool, activeOnly bool, historicalOnly bool, timeHistoricalCutoff *uint32) (*TradeOffersResult, error) {
	if !getSent && !getReceived {
		return nil, fmt.Errorf("getSent and getReceived can't be both false\n")
	}

	params := map[string]string{
		"key": string(c.key),
	}
	if getSent {
		params["get_sent_offers"] = "1"
	}
	if getReceived {
		params["get_received_offers"] = "1"
	}
	if getDescriptions {
		params["get_descriptions"] = "1"
		params["language"] = "en_us"
	}
	if activeOnly {
		params["active_only"] = "1"
	}
	if historicalOnly {
		params["historical_only"] = "1"
	}
	if timeHistoricalCutoff != nil {
		params["time_historical_cutoff"] = strconv.FormatUint(uint64(*timeHistoricalCutoff), 10)
	}
	resp, err := c.client.Get(fmt.Sprintf(apiUrl, "GetTradeOffers", 1) + "?" + netutil.ToUrlValues(params).Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	t := new(struct {
		Response *TradeOffersResult
	})
	if err = json.NewDecoder(resp.Body).Decode(t); err != nil {
		return nil, err
	}
	return t.Response, nil
}

func (c *Client) action(method string, version uint, offerId uint64) error {
	resp, err := c.client.Do(netutil.NewPostForm(fmt.Sprintf(apiUrl, method, version), netutil.ToUrlValues(map[string]string{
		"key":          string(c.key),
		"tradeofferid": strconv.FormatUint(offerId, 10),
	})))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf(method+" error: status code %d", resp.StatusCode)
	}

	var result map[string]string
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return err
	}
	success, ok := result["success"]
	if !ok {
		return newSteamErrorf(method + " error: no success field\n")
	}
	if success != "1" {
		return newSteamErrorf(method+" error: success field is %v\n", success)
	}
	return nil
}

func (c *Client) Decline(offerId uint64) error {
	return c.action("DeclineTradeOffer", 1, offerId)
}

func (c *Client) Cancel(offerId uint64) error {
	return c.action("CancelTradeOffer", 1, offerId)
}

// on success returns trade id
func (c *Client) Accept(offerId uint64) (uint64, error) {
	baseurl := fmt.Sprintf("https://steamcommunity.com/tradeoffer/%d/", offerId)
	req := netutil.NewPostForm(baseurl+"accept", netutil.ToUrlValues(map[string]string{
		"sessionid":    c.sessionId,
		"serverid":     "1",
		"tradeofferid": strconv.FormatUint(offerId, 10),
	}))
	req.Header.Add("Referer", baseurl)

	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("accept error: status code %d", resp.StatusCode)
	}
	var result map[string]string
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return 0, err
	}
	if strError, ok := result["strError"]; ok {
		return 0, newSteamErrorf("accept error: %v\n", strError)
	}
	tradeIdString, ok := result["tradeid"]
	if !ok {
		return 0, newSteamErrorf("accept error: steam does not return trade id\n")
	}
	tradeId, err := strconv.ParseUint(tradeIdString, 10, 64)
	if err != nil || tradeId == 0 {
		return 0, newSteamErrorf("accept error: steam returned %v for trade id", tradeIdString)
	}
	return tradeId, nil
}

type TradeItem struct {
	AppId      uint32
	ContextId  uint64
	Amount     uint64
	AssetId    uint64 `json:",string,omitempty"`
	CurrencyId uint64 `json:",string,omitempty"`
}

// Sends a new trade offer to the given Steam user. You can optionally specify an access token if you've got one.
// In addition, `counteredOfferId` can be non-nil, indicating the trade offer this is a counter for.
// On success returns trade offer id
func (c *Client) Create(other steamid.SteamId, accessToken *string, myItems, theirItems []TradeItem, counteredOfferId *uint64, message string) (uint64, error) {
	// Create new trade offer status
	to := map[string]interface{}{
		"newversion": true,
		"version":    3,
		"me": map[string]interface{}{
			"assets":   myItems,
			"currency": make([]struct{}, 0),
			"ready":    false,
		},
		"them": map[string]interface{}{
			"assets":   theirItems,
			"currency": make([]struct{}, 0),
			"ready":    false,
		},
	}

	jto, err := json.Marshal(to)
	if err != nil {
		panic(err)
	}

	// Create url parameters for request
	data := map[string]string{
		"sessionid":         c.sessionId,
		"serverid":          "1",
		"partner":           other.ToString(),
		"tradeoffermessage": message,
		"json_tradeoffer":   string(jto),
		"captcha":           "",
	}

	var referer string
	if counteredOfferId != nil {
		referer = fmt.Sprintf("https://steamcommunity.com/tradeoffer/%d/", *counteredOfferId)
		data["tradeofferid_countered"] = strconv.FormatUint(*counteredOfferId, 10)
	} else {
		// Add token for non-friend offers
		if accessToken != nil {
			params := map[string]string{
				"trade_offer_access_token": *accessToken,
			}
			paramsJson, err := json.Marshal(params)
			if err != nil {
				panic(err)
			}

			data["trade_offer_create_params"] = string(paramsJson)

			referer = "https://steamcommunity.com/tradeoffer/new/?partner=" + strconv.FormatUint(uint64(other.GetAccountId()), 10) + "&token=" + *accessToken
		} else {

			referer = "https://steamcommunity.com/tradeoffer/new/?partner=" + strconv.FormatUint(uint64(other.GetAccountId()), 10)
		}
	}

	// Create request
	req := netutil.NewPostForm("https://steamcommunity.com/tradeoffer/new/send", netutil.ToUrlValues(data))
	req.Header.Add("Referer", referer)

	// Send request
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, fmt.Errorf("create error: status code %d", resp.StatusCode)
	}
	var result map[string]string
	if err = json.NewDecoder(resp.Body).Decode(result); err != nil {
		return 0, err
	}
	if strError, ok := result["strError"]; ok {
		return 0, newSteamErrorf("create error: %v\n", strError)
	}
	offerIdString, ok := result["tradeofferid"]
	if !ok {
		return 0, newSteamErrorf("create error: steam does not return trade offer id\n")
	}
	offerId, err := strconv.ParseUint(offerIdString, 10, 64)
	if err != nil || offerId == 0 {
		return 0, newSteamErrorf("create error: steam returned %v for trade offer id", offerIdString)
	}
	return offerId, nil
}

func (c *Client) GetOwnInventory(contextId uint64, appId uint32) (*inventory.Inventory, error) {
	return inventory.GetOwnInventory(c.client, contextId, appId)
}

func (c *Client) GetTheirInventory(other steamid.SteamId, contextId uint64, appId uint32) (*inventory.Inventory, error) {
	return inventory.GetFullInventory(func() (*inventory.PartialInventory, error) {
		return c.getPartialTheirInventory(other, contextId, appId, nil)
	}, func(start uint) (*inventory.PartialInventory, error) {
		return c.getPartialTheirInventory(other, contextId, appId, &start)
	})
}

func (c *Client) getPartialTheirInventory(other steamid.SteamId, contextId uint64, appId uint32, start *uint) (*inventory.PartialInventory, error) {
	data := map[string]string{
		"sessionid": c.sessionId,
		"partner":   fmt.Sprintf("%d", other),
		"contextid": strconv.FormatUint(contextId, 10),
		"appid":     strconv.FormatUint(uint64(appId), 10),
	}
	if start != nil {
		data["start"] = strconv.FormatUint(uint64(*start), 10)
	}

	const baseUrl = "https://steamcommunity.com/tradeoffer/new/"
	req, err := http.NewRequest("GET", baseUrl+"partnerinventory/?"+netutil.ToUrlValues(data).Encode(), nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Referer", baseUrl+"?partner="+fmt.Sprintf("%d", other))

	return inventory.DoInventoryRequest(c.client, req)
}

func (c *Client) GetOfferWithRetry(offerId uint64, retryCount int, retryWait time.Duration) (*TradeOfferResult, error) {
	var res *TradeOfferResult
	return res, withRetry(
		func() (err error) {
			res, err = c.GetOffer(offerId)
			return err
		}, retryCount, retryWait)
}

func (c *Client) GetOffersWithRetry(getSent bool, getReceived bool, getDescriptions bool, activeOnly bool, historicalOnly bool, timeHistoricalCutoff *uint32, retryCount int, retryWait time.Duration) (*TradeOffersResult, error) {
	var res *TradeOffersResult
	return res, withRetry(
		func() (err error) {
			res, err = c.GetOffers(getSent, getReceived, getDescriptions, activeOnly, historicalOnly, timeHistoricalCutoff)
			return err
		}, retryCount, retryWait)
}

func (c *Client) DeclineWithRetry(offerId uint64, retryCount int, retryWait time.Duration) error {
	return withRetry(
		func() error {
			return c.Decline(offerId)
		}, retryCount, retryWait)
}

func (c *Client) CancelWithRetry(offerId uint64, retryCount int, retryWait time.Duration) error {
	return withRetry(
		func() error {
			return c.Cancel(offerId)
		}, retryCount, retryWait)
}

func (c *Client) AcceptWithRetry(offerId uint64, retryCount int, retryWait time.Duration) (uint64, error) {
	var res uint64
	return res, withRetry(
		func() (err error) {
			res, err = c.Accept(offerId)
			return err
		}, retryCount, retryWait)
}

func (c *Client) CreateWithRetry(other steamid.SteamId, accessToken *string, myItems, theirItems []TradeItem, counteredOfferId *uint64, message string, retryCount int, retryWait time.Duration) (uint64, error) {
	var res uint64
	return res, withRetry(
		func() (err error) {
			res, err = c.Create(other, accessToken, myItems, theirItems, counteredOfferId, message)
			return err
		}, retryCount, retryWait)
}

func (c *Client) GetOwnInventoryWithRetry(contextId uint64, appId uint32, retryCount int, retryWait time.Duration) (*inventory.Inventory, error) {
	var res *inventory.Inventory
	return res, withRetry(
		func() (err error) {
			res, err = c.GetOwnInventory(contextId, appId)
			return err
		}, retryCount, retryWait)
}

func (c *Client) GetTheirInventoryWithRetry(other steamid.SteamId, contextId uint64, appId uint32, retryCount int, retryWait time.Duration) (*inventory.Inventory, error) {
	var res *inventory.Inventory
	return res, withRetry(
		func() (err error) {
			res, err = c.GetTheirInventory(other, contextId, appId)
			return err
		}, retryCount, retryWait)
}

func withRetry(f func() error, retryCount int, retryWait time.Duration) error {
	if retryCount <= 0 {
		panic("retry count must be more than 0")
	}
	i := 0
	for {
		i++
		if err := f(); err != nil {
			// If we got steam error do not retry
			if _, ok := err.(*SteamError); ok {
				return err
			}
			if i == retryCount {
				return err
			}
			time.Sleep(retryWait)
			continue
		}
		break
	}
	return nil
}
