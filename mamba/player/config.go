package player

type PlayerConfig struct {
	ShotBase          float64
	Mamba             float64
	MentalCoefficient float64
	HealRate          float64
	ShotAccuracyRate  float64
	LogisticK         float64
	LogisticX0        float64
}

type PlayerState struct {
	Mental       float64
	ShotAccuracy float64
}

type PlayerStats struct {
	ShotsMade      int
	ShotsMissed    int
	ShotsAttempted int
}