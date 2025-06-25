package game

type StatsRecord struct {
	Poss          int     `json:"poss"`
	Score         int     `json:"score"`
	Shooter       int     `json:"shooter"`
	Mental1       float64 `json:"mental1"`
	ShotAccuracy1 float64 `json:"shot_accuracy1"`
	Made1         int     `json:"made1"`
	Missed1       int     `json:"maded1"`
	Mental2       float64 `json:"mental2"`
	ShotAccuracy2 float64 `json:"shot_accuracy2"`
	Made2         int     `json:"made2"`
	Missed2       int     `json:"missed2"`
	Mental3       float64 `json:"mental3"`
	ShotAccuracy3 float64 `json:"shot_accuracy3"`
	Made3         int     `json:"made3"`
	Missed3       int     `json:"missed3"`
	Mental4       float64 `json:"mental4"`
	ShotAccuracy4 float64 `json:"shot_accuracy4"`
	Made4         int     `json:"made4"`
	Missed4       int     `json:"missed4"`
	Mental5       float64 `json:"mental5"`
	ShotAccuracy5 float64 `json:"shot_accuracy5"`
	Made5         int     `json:"made5"`
	Missed5       int     `json:"missed5"`
}
