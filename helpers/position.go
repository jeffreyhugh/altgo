package helpers

import (
	"github.com/Tnze/go-mc/bot/world/entity/player"
	"math"
)

func OccupiesSameBlock(player1, player2 player.Pos) bool {
	return (math.Abs(player1.X - player2.X) <= 1) &&
		(math.Abs(player1.Z - player2.Z) <= 1) &&
		(math.Abs(player1.Y - player2.Y) <= 0.6)
}