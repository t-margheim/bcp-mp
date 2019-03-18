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
			name:    "invalid date -- too early",
			date:    time.Date(1985, 2, 28, 0, 0, 0, 0, time.UTC),
			want:    KeyChain{},
			wantErr: true,
		},
		{
			name: "Valid date -- Mar 17 2019",
			date: time.Date(2019, 3, 17, 0, 0, 0, 0, time.UTC),
			want: KeyChain{
				Season:    SeasonLent,
				Open:      SeasonLent,
				Week:      1,
				Weekday:   "Sunday",
				ShortDate: "Mar 17",
				Year:      1,
				Iterator:  440,
				Date:      time.Date(2019, 3, 17, 0, 0, 0, 0, time.UTC),
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
				t.Errorf("GetKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSeason(t *testing.T) {
	type args struct {
		date time.Time
	}
	tests := []struct {
		name string
		args args
		want KeyChain
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSeason(tt.args.date); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetSeason() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOpen(t *testing.T) {
	type args struct {
		date   time.Time
		season Key
	}
	tests := []struct {
		name string
		args args
		want Key
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOpen(tt.args.date, tt.args.season); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOpen() = %v, want %v", got, tt.want)
			}
		})
	}
}
