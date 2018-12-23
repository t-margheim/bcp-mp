package calendar

import (
	"reflect"
	"testing"
	"time"
)

func TestGetKeys(t *testing.T) {
	tests := []struct {
		name    string
		date    time.Time
		want    KeyChain
		wantErr bool
	}{
		{
			name: "Trinity Sunday",
			date: time.Date(2019, 6, 16, 0, 0, 0, 0, time.UTC),
			want: KeyChain{
				Season: SeasonOrdinary,
				Open:   OpenTrinitySunday,
			},
			wantErr: false,
		},
		{
			name: "Advent",
			date: time.Date(2018, 12, 16, 0, 0, 0, 0, time.UTC),
			want: KeyChain{
				Season: SeasonAdvent,
				Open:   SeasonAdvent,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetKeys(tt.date)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetKeys() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func Test_getOpen(t *testing.T) {
	tests := []struct {
		name   string
		date   time.Time
		season Key
		want   Key
	}{
		{
			name:   "Advent",
			date:   time.Date(2018, 12, 24, 0, 0, 0, 0, time.UTC),
			season: SeasonAdvent,
			want:   SeasonAdvent,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOpen(tt.date, tt.season); got != tt.want {
				t.Errorf("getOpen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSeason(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want Key
	}{
		{
			name: "Advent",
			date: time.Date(2018, 12, 16, 0, 0, 0, 0, time.UTC),
			want: SeasonAdvent,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSeason(tt.date); got != tt.want {
				t.Errorf("getSeason() = %v, want %v", got, tt.want)
			}
		})
	}
}
