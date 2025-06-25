package team

import (
	"testing"

	"github.com/poteto0/mamba/player"
)

func TestNewTeam(t *testing.T) {
	players := []player.IPlayer{
		player.NewPlayer(player.PlayerConfig{}),
		player.NewPlayer(player.PlayerConfig{}),
		player.NewPlayer(player.PlayerConfig{}),
		player.NewPlayer(player.PlayerConfig{}),
		player.NewPlayer(player.PlayerConfig{}),
	}

	team := NewTeam(players)

	if team == nil {
		t.Error("NewTeam should not return nil")
	}
}

func TestTeam_SelectShooter(t *testing.T) {
	players := []player.IPlayer{
		player.NewPlayer(player.PlayerConfig{}),
		player.NewPlayer(player.PlayerConfig{}),
		player.NewPlayer(player.PlayerConfig{}),
		player.NewPlayer(player.PlayerConfig{}),
		player.NewPlayer(player.PlayerConfig{}),
	}

	team := NewTeam(players).(*Team)

	shooterIndex := team.SelectShooter()

	if shooterIndex < 0 || shooterIndex > 5 {
		t.Errorf("Expected shooterIndex to be between 0 and 5, but got %d", shooterIndex)
	}
}
