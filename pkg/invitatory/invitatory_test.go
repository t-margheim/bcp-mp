package invitatory

import (
	"reflect"
	"testing"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name string
		keys calendar.KeyChain
		want Entry
	}{
		{
			name: "easter opening",
			keys: calendar.KeyChain{
				Season: calendar.SeasonEaster,
			},
			want: easter,
		},
		{
			name: "iterator opening",
			keys: calendar.KeyChain{
				Season:   calendar.SeasonOrdinary,
				Iterator: 3,
			},
			want: options[1],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
