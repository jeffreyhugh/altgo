package helpers

import (
	"errors"
	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/bot/world/entity/player"
	"github.com/qbxt/altgo/constants"
	"github.com/qbxt/altgo/logger"
	"github.com/qbxt/altgo/structures"
	"github.com/sirupsen/logrus"
	"time"
)

func Follow(currentSession *structures.MinecraftLogin, seconds int) {
	ticks := seconds * 20

	for i := 0; i < ticks; i++ {
		newPosition, err := getNewPosition(currentSession)
		if err != nil {
			return
		}
		if err := currentSession.Client.Physics.ServerPositionUpdate(*newPosition, &currentSession.Client.Wd); err != nil {
			logger.Error("could not change positions", err, logrus.Fields{"username": currentSession.IGN})
		}
		time.Sleep(50 * time.Millisecond) // 1 tick
	}

	logger.Info("follow finished", logrus.Fields{"username": currentSession.IGN})
}

func GoTo(currentSession *structures.MinecraftLogin) error {
	newPosition, err := getNewPosition(currentSession)
	if err != nil {
		return err
	}
	if err := currentSession.Client.Physics.ServerPositionUpdate(*newPosition, &currentSession.Client.Wd); err != nil {
		logger.Error("could not change positions", err, logrus.Fields{"username": currentSession.IGN})
	}

	return nil
}

func getNewPosition(currentSession *structures.MinecraftLogin) (*player.Pos, error) {
	players := currentSession.Client.Wd.PlayerEntities()
	foundPlayer := entity.Entity{}
	for _, p := range players {
		if p.UUID.String() == constants.UUID_QUEUEBOT {
			foundPlayer = p
		}
	}

	if foundPlayer.ID == 0 {
		return nil, errors.New("unknown player")
	}

	return &player.Pos{
		X:        foundPlayer.X,
		Y:        foundPlayer.Y,
		Z:        foundPlayer.Z,
		OnGround: foundPlayer.OnGround,
		Yaw:      float32(foundPlayer.Yaw),
		Pitch:    float32(foundPlayer.Pitch),
	}, nil
}
