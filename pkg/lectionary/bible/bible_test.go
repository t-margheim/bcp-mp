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

	tests := []struct {
		name      string
		baseURL   string
		reference string
		want      *Lesson
	}{
		{
			name:      "OK",
			baseURL:   "http://testservice.net/fakeapi?q=%s",
			reference: "Isa 40:1-11",
			want: &Lesson{
				Reference: "Isaiah 40:1-11",
				Body:      "This is just a test body.",
			},
		},
		{
			name:      "badURL",
			baseURL:   "12345",
			reference: "Isa 40:1-11",
			want: &Lesson{
				Reference: "Failed on http.NewRequest()",
				Body:      `error message: parse 12345%!(EXTRA string=Isa+40%3A1-11): invalid URL escape "%!("`,
			},
		},
		{
			name:      "badURL",
			baseURL:   "http://test-example.net/fake/%s",
			reference: "Isa 40:1-11",
			want: &Lesson{
				Reference: "Failed on client.Do()",
				Body:      `error message: Get http://test-example.net/fake/Isa+40%3A1-11: no responder found`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				BaseURL: tt.baseURL,
				client:  &http.Client{},
			}
			if got := s.GetLesson(tt.reference); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.getLesson() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
