package opening

import (
	"reflect"
	"testing"
	"time"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name    string
		date    time.Time
		want    []Opening
		wantErr bool
	}{
		{
			name:    "Trinity Sunday",
			date:    time.Date(2019, 6, 16, 0, 0, 0, 0, time.UTC),
			want:    Openings[calendar.OpenTrinitySunday],
			wantErr: false,
		},
		{
			name:    "Invalid Date",
			date:    time.Date(2011, 6, 16, 0, 0, 0, 0, time.UTC),
			want:    []Opening{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.date)

			if err != nil {
				if tt.wantErr {
					return
				}
				t.Errorf("unexpected error: %s", err.Error())
				return
			}

			var validReturn bool
			for _, o := range tt.want {
				if reflect.DeepEqual(got, o) {
					validReturn = true
					break
				}
			}
			if !validReturn {
				t.Errorf("incorrect opening returned, got %+v, wanted one of %+v", got, tt.want)
			}
		})
	}
}
