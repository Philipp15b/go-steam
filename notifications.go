package steam

import (
	. "github.com/Philipp15b/go-steam/internal"
	. "github.com/Philipp15b/go-steam/internal/protobuf"
	. "github.com/Philipp15b/go-steam/internal/steamlang"
)

type Notifications struct {
	// Maps notification types to their count. If a type is not present in the map,
	// its count is zero.
	notifications map[NotificationType]uint
	client        *Client
}

func newNotifications(client *Client) *Notifications {
	return &Notifications{
		make(map[NotificationType]uint),
		client,
	}
}

func (n *Notifications) HandlePacket(packet *Packet) {
	switch packet.EMsg {
	case EMsg_ClientUserNotifications:
		n.handleClientUserNotifications(packet)
	}
}

type NotificationType uint

const (
	TradeOffer NotificationType = 1
)

// This event is emitted for every CMsgClientUserNotifications message and likewise only used for
// trade offers. Unlike the the above it is also emitted when the count of a type that was tracked
// before by this Notifications instance reaches zero.
type NotificationEvent struct {
	Type  NotificationType
	Count uint
}

func (n *Notifications) handleClientUserNotifications(packet *Packet) {
	msg := new(CMsgClientUserNotifications)
	packet.ReadProtoMsg(msg)

	for _, notification := range msg.GetNotifications() {
		typ := NotificationType(*notification.UserNotificationType)
		count := uint(*notification.Count)
		n.notifications[typ] = count
		n.client.Emit(&NotificationEvent{typ, count})
	}

	// check if there is a notification in our map that isn't in the current packet
	for typ, _ := range n.notifications {
		exists := false
		for _, t := range msg.GetNotifications() {
			if NotificationType(*t.UserNotificationType) == typ {
				exists = true
				break
			}
		}

		if !exists {
			delete(n.notifications, typ)
			n.client.Emit(&NotificationEvent{typ, 0})
		}
	}
}
