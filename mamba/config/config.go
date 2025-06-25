package config

type ModelConfig struct {
	MambaNum             int     `json:"mamba_num"`
	MambaShotPercentage  float64 `json:"mamba_shot_percentage"`
	PlayerShotPercentage float64 `json:"player_shot_percentage"`
	MambaMental          float64 `json:"mamba_mental"`
	MentalCoefficient    float64 `json:"mental_coefficient"`
}
