package helpers

import (
	"github.com/qbxt/altgo/structures"
)

var (
	loginManagerInbox  chan *structures.MinecraftLogin
	loginManagerOutbox chan *structures.MinecraftLogin
)

func InitLoginManagerChans() {
	loginManagerInbox = make(chan *structures.MinecraftLogin)
	loginManagerOutbox = make(chan *structures.MinecraftLogin)
}

func GetLoginManagerChans() (chan *structures.MinecraftLogin, chan *structures.MinecraftLogin) {
	return loginManagerInbox, loginManagerOutbox
}
