package player

import "github.com/poteto0/mamba/utils"

type IPlayer interface {
	Update(isMade bool)
	Heal()
	Scale()
	updateState(isMade bool)
	updateStats(isMade bool)

	calcAccuracy() float64

	Mental() float64
	ShotAccuracy() float64
	ShotsMade() int
	ShotsMissed() int
}

type Player struct {
	config PlayerConfig
	state  PlayerState
	stats  PlayerStats
}

func NewPlayer(config PlayerConfig) IPlayer {
	return &Player{
		config: config,
		state:  PlayerState{},
		stats:  PlayerStats{},
	}
}

func (p *Player) calcAccuracy() float64 {
	return float64(p.stats.ShotsMade) / float64(p.stats.ShotsAttempted)
}

func (p *Player) Heal() {
	p.state.Mental += p.config.HealRate
}

func (p *Player) Scale() {
	if p.state.Mental < 0 {
		p.state.Mental = 0
	}
}

func (p *Player) Update(isMade bool) {
	p.updateStats(isMade)
	p.updateState(isMade)
}

func (p *Player) updateState(isMade bool) {
	p.updateShotAccuracy()
	if isMade {
		p.state.Mental += p.config.MentalCoefficient * p.config.ShotBase
		return
	}

	p.state.Mental -= 2 * p.config.MentalCoefficient * (1 - p.config.Mamba)
}

func (p *Player) updateShotAccuracy() {
	currentAccuracy := p.calcAccuracy()

	logi := utils.Logistic(float64(p.stats.ShotsAttempted), p.config.LogisticK, p.config.LogisticX0, 1)
	p.state.ShotAccuracy = p.config.ShotBase + p.config.ShotAccuracyRate*logi*(currentAccuracy-p.config.ShotBase)
}

func (p *Player) updateStats(isMade bool) {
	if isMade {
		p.stats.ShotsMade++
		p.stats.ShotsAttempted++
		return
	}

	p.stats.ShotsMissed++
	p.stats.ShotsAttempted++
}

func (p *Player) Mental() float64 {
	return p.state.Mental
}

func (p *Player) ShotAccuracy() float64 {
	return p.state.ShotAccuracy
}

func (p *Player) ShotsMade() int {
	return p.stats.ShotsMade
}

func (p *Player) ShotsMissed() int {
	return p.stats.ShotsMissed
}
