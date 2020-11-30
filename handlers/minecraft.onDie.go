package handlers

import (
	"github.com/qbxt/altgo/logger"
	"github.com/qbxt/altgo/structures"
	"github.com/sirupsen/logrus"
)

func OnDie(currentSession *structures.MinecraftLogin) error {
	_ = currentSession.Client.Respawn()
	logger.Info("respawned", logrus.Fields{"ign": currentSession.IGN})
	return nil
}
