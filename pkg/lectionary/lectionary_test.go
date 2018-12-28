package lectionary

import (
	"reflect"
	"testing"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
)

func TestGetReadings(t *testing.T) {
	tests := []struct {
		name string
		keys calendar.KeyChain
		want Readings
	}{
		{
			name: "December 17",
			keys: calendar.KeyChain{
				Season:    calendar.SeasonAdvent,
				Week:      3,
				Weekday:   "Monday",
				ShortDate: "Dec 17",
				Year:      1,
			},
			want: Readings{
				Psalms: []string{"41", "52"},
				First:  "Isa 8:16–9:1",
				Second: "2 Pet 1:1–11",
				Gospel: "Luke 22:39–53",
			},
		},
		{
			name: "December 24",
			keys: calendar.KeyChain{
				Season:    calendar.SeasonAdvent,
				Week:      4,
				Weekday:   "Monday",
				ShortDate: "Dec 24",
				Year:      1,
			},
			want: Readings{
				Psalms: []string{"45", "46"},
				First:  "Isa 35:1–10",
				Second: "Rev 22:12–17, 21",
				Gospel: "Luke 1:67–80",
			},
		},
		{
			name: "December 26",
			keys: calendar.KeyChain{
				Season:    calendar.SeasonChristmas,
				Week:      0,
				Weekday:   "Wednesday",
				ShortDate: "Dec 26",
				Year:      1,
			},
			want: Readings{
				Psalms: []string{"28", "30"},
				First:  "2 Chr 24:17–22",
				Second: "Acts 6:1–7",
				Gospel: "",
				Title:  "Saint Stephen, Deacon and Martyr",
			},
		},
		{
			name: "December 29",
			keys: calendar.KeyChain{
				Season:    calendar.SeasonChristmas,
				Week:      0,
				Weekday:   "Saturday",
				ShortDate: "Dec 29",
				Year:      1,
			},
			want: Readings{
				Psalms: []string{"18:1–20"},
				First:  "Isa 12:1–6",
				Second: "Rev 1:1–8",
				Gospel: "John 7:37–52",
			},
		},
		{
			name: "August 13",
			keys: calendar.KeyChain{
				Season:    calendar.SeasonOrdinary,
				Week:      9,
				Weekday:   "Tuesday",
				ShortDate: "Aug 13",
				Year:      1,
			},
			want: Readings{
				Psalms: []string{"5", "6"},
				First:  "1 Sam 15:24–35",
				Second: "Acts 9:32–43",
				Gospel: "Luke 23:56b–24:11",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := New()
			if got := svc.GetReadings(tt.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReadings() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
