package handlers

/* Debug

import (
	"github.com/qbxt/altgo/structures"
	"github.com/qbxt/altgo/logger"
	"github.com/Tnze/go-mc/data"
	"github.com/Tnze/go-mc/net/packet"
	"github.com/sirupsen/logrus"
)

func OnPacket(currentSession *structures.MinecraftLogin, p packet.Packet) (bool, error) {
	if p.ID == int32(data.ChatClientbound) || p.ID == int32(data.NameItem) {
		logger.Info("received packet", logrus.Fields{"ign": currentSession.IGN, "id": p.ID, "data": p.Data})
	}
	return false, nil
}

*/
