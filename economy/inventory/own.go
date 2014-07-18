package inventory

import (
	"net/http"
)

func GetOwnInventory(client *http.Client, contextId uint64, appId uint32, start *uint) (*inventory.PartialInventory, error) {
	// TODO: the "trading" parameter can be left off to return non-tradable items too
	url := fmt.Sprintf("http://steamcommunity.com/my/inventory/json/%d/%d?trading=1", appId, contextId)
	if start != nil {
		url += "&start=" + strconv.FormatUint(uint64(*start), 10)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	return RunInventoryRequest(req)
}
