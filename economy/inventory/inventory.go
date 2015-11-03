/*
Includes inventory types as used in the trade package
*/
package inventory

import (
	"bytes"
	"encoding/json"
	"github.com/Philipp15b/go-steam/economy"
	"github.com/Philipp15b/go-steam/jsont"
)

type Inventory struct {
	Items        Items        `json:"rgInventory"`
	Currencies   Currencies   `json:"rgCurrency"`
	Descriptions Descriptions `json:"rgDescriptions"`
	AppInfo      *AppInfo     `json:"rgAppInfo"`
}

type Items map[string]*Item

func (i *Items) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("[]")) {
		return nil
	}
	return json.Unmarshal(data, (*map[string]*Item)(i))
}

type Currencies map[string]*Currency

func (c *Currencies) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("[]")) {
		return nil
	}
	return json.Unmarshal(data, (*map[string]*Currency)(c))
}

type Descriptions map[string]*Description

func (d *Descriptions) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("[]")) {
		return nil
	}
	return json.Unmarshal(data, (*map[string]*Description)(d))
}

type Item struct {
	Id         economy.AssetId    `json:",string"`
	ClassId    economy.ClassId    `json:",string"`
	InstanceId economy.InstanceId `json:",string"`
	Amount     uint64             `json:",string"`
	Pos        uint32
}

type Currency struct {
	Id         economy.AssetId `json:",string"`
	ClassId    economy.ClassId `json:",string"`
	IsCurrency bool            `json:"is_currency"`
	Pos        uint32
}

type Description struct {
	AppId      uint32             `json:",string"`
	ClassId    economy.ClassId    `json:",string"`
	InstanceId economy.InstanceId `json:",string"`

	IconUrl      string `json:"icon_url"`
	IconUrlLarge string `json:"icon_url_large"`
	IconDragUrl  string `json:"icon_drag_url"`

	Name           string
	MarketName     string `json:"market_name"`
	MarketHashName string `json:"market_hash_name"`

	// Colors in hex, for example `B2B2B2`
	NameColor       string `json:"name_color"`
	BackgroundColor string `json:"background_color"`

	Type string

	Tradable   jsont.UintBool
	Marketable jsont.UintBool
	Commodity  jsont.UintBool

	Descriptions DescriptionLines
	Actions      []*Action
	// Application-specific data, like "def_index" and "quality" for TF2
	AppData map[string]string
}

type DescriptionLines []*DescriptionLine

func (d *DescriptionLines) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte(`""`)) {
		return nil
	}
	return json.Unmarshal(data, (*[]*DescriptionLine)(d))
}

type DescriptionLine struct {
	Value string
	Type  *string // Is `html` for HTML descriptions
	Color *string
}

type Action struct {
	Name string
	Link string
}

type AppInfo struct {
	AppId uint32
	Name  string
	Icon  string
	Link  string
}
