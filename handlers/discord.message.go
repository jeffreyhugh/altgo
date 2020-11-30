package handlers

import (
	"github.com/qbxt/altgo/constants"
	"github.com/qbxt/altgo/logger"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"strings"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID { // do not respond to self
		return
	}

	app, err := s.Application("@me")
	if err != nil {
		logger.Error("could not get application", err, logrus.Fields{"appid": s.State.User.ID})
		return
	}

	if m.Author.ID != app.Owner.ID { // Owner only, remove/change this check if necessary
		return
	}

	if strings.HasPrefix(m.Content, constants.DISCORD_PREFIX) {
		if strings.HasPrefix(m.Content, fmt.Sprintf("%sjoin", constants.DISCORD_PREFIX)) || strings.HasPrefix(m.Content, fmt.Sprintf("%sconnect", constants.DISCORD_PREFIX)){
			ConnectCommand(s, m)
		} else if strings.HasPrefix(m.Content, fmt.Sprintf("%sdisconnect", constants.DISCORD_PREFIX)) {
			DisconnectCommand(s, m)
		} else if strings.HasPrefix(m.Content, fmt.Sprintf("%schat", constants.DISCORD_PREFIX)) {
			ChatCommand(s, m)
		}
	}
}
