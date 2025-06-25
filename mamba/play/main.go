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
	file, err := os.Open("../_fixture/sample_config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	file.Read(buffer)

	config := config.ModelConfig{}
	if err := json.Unmarshal(buffer, &config); err != nil {
		log.Fatal(err)
	}

	players := generator.Generate(config)
	game := game.NewGame(players)

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
