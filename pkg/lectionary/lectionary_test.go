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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetReadings(tt.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetReadings() = %v, want %v", got, tt.want)
			}
		})
	}
}
