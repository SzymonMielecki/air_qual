package utils

import "time"

type AutoGenerated struct {
	Status string `json:"status"`
	Data   struct {
		City     string `json:"city"`
		State    string `json:"state"`
		Country  string `json:"country"`
		Location struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"location"`
		Current struct {
			Pollution struct {
				Ts     time.Time `json:"ts"`
				Mainus string    `json:"mainus"`
				Maincn string    `json:"maincn"`
				Aqius  int       `json:"aqius"`
				Aqicn  int       `json:"aqicn"`
			} `json:"pollution"`
			Weather struct {
				Ts time.Time `json:"ts"`
				Ic string    `json:"ic"`
				Tp int       `json:"tp"`
				Pr int       `json:"pr"`
				Hu int       `json:"hu"`
				Ws float64   `json:"ws"`
				Wd int       `json:"wd"`
			} `json:"weather"`
		} `json:"current"`
	} `json:"data"`
}

type Pollution struct {
	Maincn string
	Mainus string
	Aqius  int
	Aqicn  int
}

type Weather struct {
	Temperature   int
	Pressure      int
	Humidity      int
	WindSpeed     float64
	WindDirection int
}
