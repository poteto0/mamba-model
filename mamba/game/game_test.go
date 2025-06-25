package game

import (
	"testing"

	"github.com/poteto0/mamba/player"
)

func TestNewGame(t *testing.T) {
	players := []player.IPlayer{
		player.NewPlayer(player.PlayerConfig{}),
		player.NewPlayer(player.PlayerConfig{}),
		player.NewPlayer(player.PlayerConfig{}),
		player.NewPlayer(player.PlayerConfig{}),
		player.NewPlayer(player.PlayerConfig{}),
	}

	config := GameConfig{
		FallbackShotPercentage: 0.3,
	}

	game := NewGame(players, config)

	if game == nil {
		t.Error("NewGame should not return nil")
	}
}

func TestGame_Play(t *testing.T) {
	players := []player.IPlayer{
		player.NewPlayer(player.PlayerConfig{ShotBase: 1}),
		player.NewPlayer(player.PlayerConfig{ShotBase: 1}),
		player.NewPlayer(player.PlayerConfig{ShotBase: 1}),
		player.NewPlayer(player.PlayerConfig{ShotBase: 1}),
		player.NewPlayer(player.PlayerConfig{ShotBase: 1}),
	}

	config := GameConfig{
		FallbackShotPercentage: 0.3,
	}

	game := NewGame(players, config).(*Game)

	game.Play()

	if game.poss != MaxPossession {
		t.Errorf("Expected poss to be %d, but got %d", MaxPossession, game.poss)
	}
}
