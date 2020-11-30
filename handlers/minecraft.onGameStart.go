package handlers

import (
	"github.com/qbxt/altgo/logger"
	"github.com/qbxt/altgo/structures"
)

func OnGameStart(currentSession *structures.MinecraftLogin) error {
	logger.Info("game started", nil)

	return nil
}
