package generator

import (
	"github.com/poteto0/mamba/config"
	"github.com/poteto0/mamba/player"
)

func Generate(config config.ModelConfig) []player.IPlayer {
	players := make([]player.IPlayer, 5)

	for i := range config.MambaNum {
		mamba := player.NewPlayer(
			player.PlayerConfig{
				ShotBase:          config.MambaShotPercentage,
				Mamba:             config.MambaMental,
				MentalCoefficient: config.MentalCoefficient,
			},
		)
		players[i] = mamba
	}

	for i := config.MambaNum; i < len(players); i++ {
		playerMental := (2.5 - float64(config.MambaNum)*config.MambaMental) / (5 - float64(config.MambaNum))
		player := player.NewPlayer(
			player.PlayerConfig{
				ShotBase:          config.PlayerShotPercentage,
				Mamba:             playerMental,
				MentalCoefficient: config.MentalCoefficient,
			},
		)
		players[i] = player
	}

	return players
}
