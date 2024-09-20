package logic

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/SzymonMielecki/GoRestDemo/server/persistance"
	"github.com/SzymonMielecki/GoRestDemo/server/utils"
)

type AppState struct {
	db *persistance.Db
}

func NewAppState() *AppState {
	return &AppState{
		db: persistance.NewDb(),
	}
}

func (a *AppState) AddPollution(t time.Time, p utils.Pollution) {
	if p.Aqius < 0 {
		return
	}
	if p.Aqicn < 0 {
		return
	}

	a.db.AddPollution(p, t)
}

func (a *AppState) AddWeather(t time.Time, w utils.Weather) {
	if w.Temperature < -273 {
		return
	}
	if w.Pressure < 0 {
		return
	}
	if w.Humidity < 0 || w.Humidity > 100 {
		return
	}
	if w.WindSpeed < 0 {
		return
	}
	if w.WindDirection < 0 || w.WindDirection > 360 {
		return
	}

	a.db.AddWeather(w, t)
}

func (a *AppState) AddUniversal(data []byte) {
	auto := utils.AutoGenerated{}
	if err := json.Unmarshal(data, &auto); err != nil {
		fmt.Println("Error: ", err)
		return
	}

	a.AddPollution(auto.Data.Current.Pollution.Ts, utils.Pollution{
		Maincn: auto.Data.Current.Pollution.Maincn,
		Mainus: auto.Data.Current.Pollution.Mainus,
		Aqius:  auto.Data.Current.Pollution.Aqius,
		Aqicn:  auto.Data.Current.Pollution.Aqicn,
	})

	a.AddWeather(auto.Data.Current.Weather.Ts, utils.Weather{
		Temperature:   auto.Data.Current.Weather.Tp,
		Pressure:      auto.Data.Current.Weather.Pr,
		Humidity:      auto.Data.Current.Weather.Hu,
		WindSpeed:     auto.Data.Current.Weather.Ws,
		WindDirection: auto.Data.Current.Weather.Wd,
	})
}

func (a *AppState) GetClosestPolution(t time.Time) (utils.Pollution, error) {
	return a.db.GetClosestPolution(t)
}

func (a *AppState) GetClosestWeather(t time.Time) (utils.Weather, error) {
	return a.db.GetClosestWeather(t)
}
