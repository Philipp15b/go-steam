package inventory

import (
	"bytes"
	"encoding/json"
)

// A partial inventory as sent by the Steam API.
type PartialInventory struct {
	Success bool
	Inventory
	More      bool
	MoreStart MoreStart `json:"more_start"`
}

type MoreStart uint

func (m *MoreStart) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("false")) {
		return nil
	}
	return json.Unmarshal(data, (*uint)(m))
}

// Merges the given Inventory into a single Inventory.
// The given slice must have at least one element. The first element of the slice is used
// and modified.
func Merge(p ...*Inventory) *Inventory {
	inv := p[0]
	for idx, i := range p {
		if idx == 0 {
			continue
		}

		for key, value := range i.Items {
			inv.Items[key] = value
		}
		for key, value := range i.Descriptions {
			inv.Descriptions[key] = value
		}
		for key, value := range i.Currencies {
			inv.Currencies[key] = value
		}
	}

	return inv
}
