package handlers

import (
	"github.com/Tnze/go-mc/chat"
	"github.com/qbxt/altgo/helpers"
	"github.com/qbxt/altgo/logger"
	"github.com/qbxt/altgo/structures"
	"github.com/sirupsen/logrus"
)

func OnDisconnect(currentSession *structures.MinecraftLogin, msg chat.Message) error {
	inbox, outbox := helpers.GetLoginManagerChans()

	// Clear session when disconnected
	currentSession.Client = nil
	currentSession.Server.ConnectedAt = nil
	currentSession.Server.Hostname = ""
	currentSession.Server.Port = 0
	currentSession.Password = "a"

	inbox <- currentSession
	_ = <-outbox

	logger.Info("disconnected", logrus.Fields{"message": msg, "ign": currentSession.IGN})
	return nil
}
