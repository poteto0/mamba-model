package generator

import (
	"github.com/poteto0/mamba/player"
)

type Config struct {
	MambaNum             int
	MambaShotPercentage  float64
	PlayerShotPercentage float64
	MambaMental          float64
	MentalCoefficient    float64
	HealRate             float64
	ShotAccuracyRate     float64
	LogisticK            float64
	LogisticX0           float64
}

func Generate(config Config) []player.IPlayer {
	players := make([]player.IPlayer, 5)

	for i := 0; i < config.MambaNum; i++ {
		mamba := player.NewPlayer(
			player.PlayerConfig{
				ShotBase:          config.MambaShotPercentage,
				Mamba:             config.MambaMental,
				MentalCoefficient: config.MentalCoefficient,
				HealRate:          config.HealRate,
				ShotAccuracyRate:  config.ShotAccuracyRate,
				LogisticK:         config.LogisticK,
				LogisticX0:        config.LogisticX0,
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
				HealRate:          config.HealRate,
				ShotAccuracyRate:  config.ShotAccuracyRate,
				LogisticK:         config.LogisticK,
				LogisticX0:        config.LogisticX0,
			},
		)
		players[i] = player
	}

	return players
}