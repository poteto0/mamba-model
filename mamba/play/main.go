package main

import (
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/poteto0/mamba/config"
	"github.com/poteto0/mamba/game"
	"github.com/poteto0/mamba/generator"
)

func main() {
	config, err := config.Load("../_fixture/sample_config.json")
	if err != nil {
		log.Fatal(err)
	}

		players := generator.Generate(generator.Config{
		MambaNum:             config.MambaNum,
		MambaShotPercentage:  config.MambaShotPercentage,
		PlayerShotPercentage: config.PlayerShotPercentage,
		MambaMental:          config.MambaMental,
		MentalCoefficient:    config.MentalCoefficient,
		HealRate:             config.HealRate,
		ShotAccuracyRate:     config.ShotAccuracyRate,
		LogisticK:            config.LogisticK,
		LogisticX0:           config.LogisticX0,
	})
	game := game.NewGame(players, game.GameConfig{
		FallbackShotPercentage: config.FallbackShotPercentage,
	})

	game.Play()

	statsSheet := game.StatsSheet()
	bytes, err := json.Marshal(statsSheet)
	if err != nil {
		log.Fatal(err)
	}

	of, err := os.Create("../_fixture/sample_output.json")
	if err != nil {
		log.Fatalf("Failed to create file: %s", err)
	}
	defer of.Close()

	_, err = of.Write(bytes)
	if err != nil {
		log.Fatalf("Failed to write to file: %s", err)
	}
}
