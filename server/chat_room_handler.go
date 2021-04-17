package main 

import (
	"github.com/timaussie/chat/client"
)

type chat_room struct {
	name string,
	participants map[*Client]bool
}

func newChatRoom(name string) * chat_room {
	return &chat_room {
		name: name,
		participants: make(map[*Client]bool)
	}
}

func (cr * chat_room) broadcast(string message) {
	for pnt := range cr.participants {
		
	}
}