package config

import (
	"os"

	"github.com/goccy/go-json"
)

type ModelConfig struct {
	MambaNum             int     `json:"mamba_num"`
	MambaShotPercentage  float64 `json:"mamba_shot_percentage"`
	PlayerShotPercentage float64 `json:"player_shot_percentage"`
	MambaMental          float64 `json:"mamba_mental"`
	MentalCoefficient    float64 `json:"mental_coefficient"`
	HealRate             float64 `json:"heal_rate"`
	ShotAccuracyRate     float64 `json:"shot_accuracy_rate"`
	LogisticK            float64 `json:"logistic_k"`
	LogisticX0           float64 `json:"logistic_x0"`
	FallbackShotPercentage float64 `json:"fallback_shot_percentage"`
}

func Load(path string) (*ModelConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buffer := make([]byte, 1024)
	n, err := file.Read(buffer)
	if err != nil {
		return nil, err
	}

	config := &ModelConfig{}
	if err := json.Unmarshal(buffer[:n], config); err != nil {
		return nil, err
	}

	return config, nil
}