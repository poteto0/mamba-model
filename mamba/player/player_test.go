package player

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	config := PlayerConfig{
		ShotBase:          0.5,
		Mamba:             0.8,
		MentalCoefficient: 0.1,
	}

	player := NewPlayer(config)

	if player == nil {
		t.Error("NewPlayer should not return nil")
	}
}

func TestPlayer_Update(t *testing.T) {
	config := PlayerConfig{
		ShotBase:          0.5,
		Mamba:             0.8,
		MentalCoefficient: 0.1,
	}

	player := NewPlayer(config).(*Player)

	player.Update(true)

	if player.stats.ShotsMade != 1 {
		t.Errorf("Expected ShotsMade to be 1, but got %d", player.stats.ShotsMade)
	}

	if player.stats.ShotsAttempted != 1 {
		t.Errorf("Expected ShotsAttempted to be 1, but got %d", player.stats.ShotsAttempted)
	}

	player.Update(false)

	if player.stats.ShotsMissed != 1 {
		t.Errorf("Expected ShotsMissed to be 1, but got %d", player.stats.ShotsMissed)
	}

	if player.stats.ShotsAttempted != 2 {
		t.Errorf("Expected ShotsAttempted to be 2, but got %d", player.stats.ShotsAttempted)
	}
}

func TestPlayer_Heal(t *testing.T) {
	config := PlayerConfig{
		HealRate: 0.1,
	}

	player := NewPlayer(config).(*Player)
	initialMental := player.state.Mental

	player.Heal()

	if player.state.Mental != initialMental+config.HealRate {
		t.Errorf("Expected Mental to be %f, but got %f", initialMental+config.HealRate, player.state.Mental)
	}
}

func TestPlayer_Scale(t *testing.T) {
	player := NewPlayer(PlayerConfig{}).(*Player)
	player.state.Mental = -0.5

	player.Scale()

	if player.state.Mental != 0 {
		t.Errorf("Expected Mental to be 0, but got %f", player.state.Mental)
	}
}
