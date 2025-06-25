package team

import (
	"github.com/poteto-go/tslice"
	"github.com/poteto0/mamba/player"
	"github.com/poteto0/mamba/utils"
)

type ITeam interface {
	SelectShooter() int
	Update(shooterIndex int, isMade bool)
}

type Team struct {
	Players []player.IPlayer
}

func NewTeam(players []player.IPlayer) ITeam {
	return &Team{
		Players: players,
	}
}

func (t *Team) SelectShooter() int {
	totalMental := tslice.Reduce(
		t.Players,
		func(acc float64, player player.IPlayer) float64 {
			return acc + player.Mental()
		},
		0,
	)

	sumRate := 0.0

	rand := utils.RandFloat64()
	for i := range 5 {
		sumRate += utils.BradleyTerry(totalMental, t.Players[i].Mental(), 0.5)

		if rand <= sumRate {
			return i
		}
	}

	// shooterの選定が行われない場合。誰も打たない場合
	return 5
}

func (t *Team) Update(shooterIndex int, isMade bool) {
	for i := range t.Players {
		if i == shooterIndex {
			t.Players[i].Update(isMade)
			t.Players[i].Scale()
			continue
		}

		t.Players[i].Heal()
		t.Players[i].Scale()
	}
}
