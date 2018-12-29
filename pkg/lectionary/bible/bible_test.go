package bible

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestService_getLesson(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respOK, _ := httpmock.NewJsonResponder(200, resp{
		Canonical: "Isaiah 40:1-11",
		Passages: []string{
			"This is just a test body.",
		},
	})
	httpmock.RegisterResponder(http.MethodGet, "http://testservice.net/fakeapi", respOK)

	type fields struct {
		// dailyOffice     map[int][]storedReadings
		// specialReadings map[string]storedReadings
		baseURL string
	}
	type args struct {
		reference string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Lesson
	}{
		{
			name: "OK",
			fields: fields{
				baseURL: "http://testservice.net/fakeapi?q=%s",
			},
			args: args{
				reference: "Isa 40:1-11",
			},
			want: Lesson{
				Reference: "Isaiah 40:1-11",
				Body:      "This is just a test body.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				BaseURL: tt.fields.baseURL,
			}
			if got := s.GetLesson(tt.args.reference); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.getLesson() = %v, want %v", got, tt.want)
			}
		})
	}
}
