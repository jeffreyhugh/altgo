package structures

import (
	mc "github.com/Tnze/go-mc/bot"
	"time"
)

const (
	MinecraftCommandFollow    = 0
	MinecraftCommandTextEntry = 1
)

type MinecraftCommand struct {
	CommandType int
	Follow      string
	TextEntry   string
}

type MinecraftServer struct {
	Hostname    string
	Port        int
	ConnectedAt *time.Time
}

type MinecraftLogin struct {
	IGN, Email, Password string
	Migrated             bool
	Client               *mc.Client
	Server               MinecraftServer
}
