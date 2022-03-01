package socialcache

import (
	"errors"
	"sync"

	"github.com/Philipp15b/go-steam/v3/steamid"
)

// nickname list is a thread safe map
type NicknamesList struct {
	mutex sync.RWMutex
	byId  map[steamid.SteamId]*Nickname
}

// Returns a new nickname list
func NewNicknamesList() *NicknamesList {
	return &NicknamesList{byId: make(map[steamid.SteamId]*Nickname)}
}

// Adds a nickname to the nickname list
func (list *NicknamesList) Add(nickname Nickname) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	_, exists := list.byId[nickname.SteamId]
	if !exists { // make sure this doesnt already exist
		list.byId[nickname.SteamId] = &nickname
	}
}

// Removes a friend from the friend list
func (list *NicknamesList) Remove(id steamid.SteamId) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	delete(list.byId, id)
}

// Returns a copy of the friends map
func (list *NicknamesList) GetCopy() map[steamid.SteamId]Nickname {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	flist := make(map[steamid.SteamId]Nickname)
	for key, friend := range list.byId {
		flist[key] = *friend
	}
	return flist
}

// Returns a copy of the friend of a given SteamId
func (list *NicknamesList) ById(id steamid.SteamId) (Nickname, error) {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	if val, ok := list.byId[id]; ok {
		return *val, nil
	}
	return Nickname{}, errors.New("Friend not found")
}

// Returns the number of friends
func (list *NicknamesList) Count() int {
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	return len(list.byId)
}

// Setter methods
func (list *NicknamesList) SetName(id steamid.SteamId, name string) {
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if val, ok := list.byId[id]; ok {
		val.Nickname = name
	}
}

// A Nickname
type Nickname struct {
	SteamId  steamid.SteamId `json:",string"`
	Nickname string
}
