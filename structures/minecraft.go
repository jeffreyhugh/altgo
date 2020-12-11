package structures

import (
	mc "github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/yggdrasil"
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
	Migrated, Write      bool
	Client               *mc.Client
	Server               MinecraftServer
	Auth                 *yggdrasil.Access
	Following            *entity.Entity
}
