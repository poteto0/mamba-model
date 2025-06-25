package game

import (
	"github.com/poteto0/mamba/player"
	"github.com/poteto0/mamba/team"
	"github.com/poteto0/mamba/utils"
)

var MaxPossession = 100

type IGame interface {
	Play()
	playPoss()
	shoot(shooterIndex int) bool
	recordStats(shooterIndex int)

	StatsSheet() []StatsRecord
}

type Game struct {
	team       *team.Team
	score      int
	poss       int
	statsSheet []StatsRecord
	config     GameConfig
}

func NewGame(players []player.IPlayer, config GameConfig) IGame {
	return &Game{
		team:       team.NewTeam(players).(*team.Team),
		score:      0,
		poss:       0,
		statsSheet: []StatsRecord{},
		config:     config,
	}
}

func (g *Game) Play() {
	for g.poss < MaxPossession {
		g.playPoss()
	}
}

func (g *Game) playPoss() {
	shooterIndex := g.team.SelectShooter()
	isMade := g.shoot(shooterIndex)
	g.team.Update(shooterIndex, isMade)

	if isMade {
		g.score += 2
	}

	g.poss++

	g.recordStats(shooterIndex)
}

func (g *Game) shoot(shooterIndex int) bool {
	rand := utils.RandFloat64()

	// 打つ人が定まらなかった時
	if shooterIndex == 5 {
		return rand <= g.config.FallbackShotPercentage
	}

	return rand <= g.team.Players[shooterIndex].ShotAccuracy()
}

func (g *Game) recordStats(shooterIndex int) {
	g.statsSheet = append(g.statsSheet, StatsRecord{
		Poss:          g.poss,
		Score:         g.score,
		Shooter:       shooterIndex + 1,
		Mental1:       g.team.Players[0].Mental(),
		ShotAccuracy1: g.team.Players[0].ShotAccuracy(),
		Made1:         g.team.Players[0].ShotsMade(),
		Missed1:       g.team.Players[0].ShotsMissed(),
		Mental2:       g.team.Players[1].Mental(),
		ShotAccuracy2: g.team.Players[1].ShotAccuracy(),
		Made2:         g.team.Players[1].ShotsMade(),
		Missed2:       g.team.Players[1].ShotsMissed(),
		Mental3:       g.team.Players[2].Mental(),
		ShotAccuracy3: g.team.Players[2].ShotAccuracy(),
		Made3:         g.team.Players[2].ShotsMade(),
		Missed3:       g.team.Players[2].ShotsMissed(),
		Mental4:       g.team.Players[3].Mental(),
		ShotAccuracy4: g.team.Players[3].ShotAccuracy(),
		Made4:         g.team.Players[3].ShotsMade(),
		Missed4:       g.team.Players[3].ShotsMissed(),
		Mental5:       g.team.Players[4].Mental(),
		ShotAccuracy5: g.team.Players[4].ShotAccuracy(),
		Made5:         g.team.Players[4].ShotsMade(),
		Missed5:       g.team.Players[4].ShotsMissed(),
	})
}

func (g *Game) StatsSheet() []StatsRecord {
	return g.statsSheet
}
