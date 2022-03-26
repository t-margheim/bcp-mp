package calendar

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
	tests := []struct {
		name    string
		date    time.Time
		want    KeyChain
		wantErr bool
	}{
		{
			name: "March 17, 2019",
			date: time.Date(2019, 3, 17, 0, 0, 0, 0, time.UTC),
			want: KeyChain{
				Season: SeasonLent,
				Week:   1,
				Year:   1,
				Date:   time.Date(2019, 3, 17, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "June 17, 2019",
			date: time.Date(2019, 6, 17, 0, 0, 0, 0, time.UTC),
			want: KeyChain{
				Season: SeasonOrdinary,
				Week:   2,
				Year:   1,
				Date:   time.Date(2019, 6, 17, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "December 17, 2018",
			date: time.Date(2018, 12, 17, 0, 0, 0, 0, time.UTC),
			want: KeyChain{
				Season: SeasonAdvent,
				Week:   3,
				Year:   1,
				Date:   time.Date(2018, 12, 17, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name: "December 17, 2019",
			date: time.Date(2019, 12, 17, 0, 0, 0, 0, time.UTC),
			want: KeyChain{
				Season: SeasonAdvent,
				Week:   3,
				Year:   0,
				Date:   time.Date(2019, 12, 17, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
		{
			name:    "February 28, 1985",
			date:    time.Date(1985, 2, 28, 0, 0, 0, 0, time.UTC),
			want:    KeyChain{},
			wantErr: true,
		},
		{
			name: "March 26, 2022",
			date: time.Date(2022, 3, 26, 0, 0, 0, 0, time.UTC),
			want: KeyChain{
				Season:    SeasonLent,
				Open:      0,
				Week:      3,
				Weekday:   "",
				ShortDate: "",
				Year:      0,
				Iterator:  0,
				Date:      time.Date(2022, 3, 26, 0, 0, 0, 0, time.UTC),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getSeason(tt.date)

			if (err != nil) != tt.wantErr {
				t.Errorf("getSeason() unexpected error state = %v, want %v", err, tt.wantErr)
			}
			assert.Equalf(t, tt.want, got, "unexpected output")
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
		{
			name: "standard day",
			args: args{
				date:   time.Date(2019, 3, 17, 0, 0, 0, 0, time.UTC),
				season: SeasonLent,
			},
			want: SeasonLent,
		},
		{
			name: "special day",
			args: args{
				date:   time.Date(2019, 6, 16, 0, 0, 0, 0, time.UTC),
				season: SeasonOrdinary,
			},
			want: OpenTrinitySunday,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOpen(tt.args.date, tt.args.season); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOpen() = %v, want %v", got, tt.want)
			}
		})
	}
}
