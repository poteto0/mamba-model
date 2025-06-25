package generator

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	config := Config{
		MambaNum:             1,
		MambaShotPercentage:  0.65,
		PlayerShotPercentage: 0.5,
		MambaMental:          1,
		MentalCoefficient:    0.1,
		HealRate:             0.05,
		ShotAccuracyRate:     0.5,
		LogisticK:            5,
		LogisticX0:           1,
	}

	players := Generate(config)

	if len(players) != 5 {
		t.Errorf("Expected 5 players, but got %d", len(players))
	}
}
