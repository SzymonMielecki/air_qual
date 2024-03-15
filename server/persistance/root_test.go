package persistance

import (
	"reflect"
	"testing"
	"time"
)

func TestNewDb(t *testing.T) {
	tests := []struct {
		name string
		want *Db
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDb(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDb_GetClosestPolution(t *testing.T) {
	type fields struct {
		pollution map[time.Time]Pollution
		weather   map[time.Time]Weather
	}
	type args struct {
		closest time.Time
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantBest Pollution
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Db{
				pollution: tt.fields.pollution,
				weather:   tt.fields.weather,
			}
			gotBest, err := d.GetClosestPolution(tt.args.closest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetClosestPolution() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBest, tt.wantBest) {
				t.Errorf("Db.GetClosestPolution() = %v, want %v", gotBest, tt.wantBest)
			}
		})
	}
}

func TestDb_GetClosestWeather(t *testing.T) {
	type fields struct {
		pollution map[time.Time]Pollution
		weather   map[time.Time]Weather
	}
	type args struct {
		closest time.Time
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantBest Weather
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Db{
				pollution: tt.fields.pollution,
				weather:   tt.fields.weather,
			}
			gotBest, err := d.GetClosestWeather(tt.args.closest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Db.GetClosestWeather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBest, tt.wantBest) {
				t.Errorf("Db.GetClosestWeather() = %v, want %v", gotBest, tt.wantBest)
			}
		})
	}
}

func TestDb_AddPollution(t *testing.T) {
	type fields struct {
		pollution map[time.Time]Pollution
		weather   map[time.Time]Weather
	}
	type args struct {
		p Pollution
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Db{
				pollution: tt.fields.pollution,
				weather:   tt.fields.weather,
			}
			d.AddPollution(tt.args.p, tt.args.t)
		})
	}
}

func TestDb_AddWeather(t *testing.T) {
	type fields struct {
		pollution map[time.Time]Pollution
		weather   map[time.Time]Weather
	}
	type args struct {
		w Weather
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Db{
				pollution: tt.fields.pollution,
				weather:   tt.fields.weather,
			}
			d.AddWeather(tt.args.w, tt.args.t)
		})
	}
}