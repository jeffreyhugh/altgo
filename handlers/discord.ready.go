package handlers

import (
	"github.com/qbxt/altgo/logger"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	logger.Info("bot is ready", logrus.Fields{"name": s.State.User.Username, "discriminator": s.State.User.Discriminator})
}
