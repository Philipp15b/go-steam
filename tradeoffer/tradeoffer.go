/*
Implements methods to interact with the official Trade Offer API.

See: https://developer.valvesoftware.com/wiki/Steam_Web_API/IEconService
*/
package tradeoffer

import (
	"github.com/Philipp15b/go-steam/economy/inventory"
	"github.com/Philipp15b/go-steam/steamid"
)

type TradeOfferState uint

const (
	TradeOfferState_Invalid       TradeOfferState = 1  // Invalid
	TradeOfferState_Active                        = 2  // This trade offer has been sent, neither party has acted on it yet.
	TradeOfferState_Accepted                      = 3  // The trade offer was accepted by the recipient and items were exchanged.
	TradeOfferState_Countered                     = 4  // The recipient made a counter offer
	TradeOfferState_Expired                       = 5  // The trade offer was not accepted before the expiration date
	TradeOfferState_Canceled                      = 6  // The sender cancelled the offer
	TradeOfferState_Declined                      = 7  // The recipient declined the offer
	TradeOfferState_InvalidItems                  = 8  // Some of the items in the offer are no longer available (indicated by the missing flag in the output)
	TradeOfferState_EmailPending                  = 9  // The offer hasn't been sent yet and is awaiting email confirmation. The offer is only visible to the sender.
	TradeOfferState_EmailCanceled                 = 10 // Either party canceled the offer via email. The offer is visible to both parties, even if the sender canceled it before it was sent.
)

type Asset struct {
	AppId      uint64 `json:",string"`
	ContextId  uint64 `json:",string"`
	AssetId    uint64 `json:",string"`
	CurrencyId uint64 `json:",string"`
	ClassId    uint64 `json:",string"`
	InstanceId uint64 `json:",string"`
	Amount     uint64 `json:",string"`
	Missing    bool
}

type TradeOffer struct {
	TradeOfferId   uint64          `json:",string"`
	OtherAccountId steamid.SteamId `json:"accountid_other"`
	Message        string
	ExpirationTime uint32          `json:"expiraton_time"`
	State          TradeOfferState `json:"trade_offer_state"`
	ToGive         []*Asset        `json:"items_to_give"`
	ToReceive      []*Asset        `json:"items_to_receive"`
	IsOurs         bool            `json:"is_our_offer"`
	TimeCreated    uint32          `json:"time_created"`
	TimeUpdated    uint32          `json:"time_updated"`
}

type TradeOffersResult struct {
	Sent         []*TradeOffer `json:"trade_offers_sent"`
	Received     []*TradeOffer `json:"trade_offers_received"`
	Descriptions inventory.Descriptions
}

type TradeOfferResult struct {
	Offer        *TradeOffer
	Descriptions inventory.Descriptions
}
