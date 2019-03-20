package main

import (
	"context"
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
	"github.com/t-margheim/bcp-mp/pkg/lectionary"
	"github.com/t-margheim/bcp-mp/pkg/lectionary/bible"
)

type mockLectionary struct {
	MockGetReadings func(context.Context, calendar.KeyChain) lectionary.Readings
}

func (s *mockLectionary) GetReadings(ctx context.Context, keys calendar.KeyChain) lectionary.Readings {
	if s.MockGetReadings != nil {
		return s.MockGetReadings(ctx, keys)
	}
	return lectionary.Readings{}
}

func Test_prayerApp_ServeHTTP(t *testing.T) {
	tests := []struct {
		name       string
		keygen     func(time.Time) (calendar.KeyChain, error)
		req        *http.Request
		lectionary *mockLectionary
		wantHTML   []string
		wantStatus int
	}{
		{
			name: "Test 1",
			keygen: func(time.Time) (calendar.KeyChain, error) {
				return calendar.KeyChain{}, nil
			},
			req: httptest.NewRequest("GET", "http://testaddress/?date=2018-12-26", nil),
			lectionary: &mockLectionary{
				MockGetReadings: func(context.Context, calendar.KeyChain) lectionary.Readings {
					return lectionary.Readings{
						First: bible.Lesson{
							Reference: "OT Reading 1:1",
							Body:      "Old Testament Lesson",
						},
						Second: bible.Lesson{
							Reference: "NT Reading 1:1",
							Body:      "New Testament Lesson",
						},
						Gospel: bible.Lesson{
							Reference: "Gospel Reading 1:1",
							Body:      "Gospel Lesson",
						},
						Psalms: bible.Lesson{
							Reference: "Psalm Readings",
							Body:      "Psalm Lesson",
						},
						Title: "St Hedwig",
					}
				},
			},
			wantHTML: []string{
				"<h1>Morning Prayer for December 26, 2018</h1>",
				"<h2>St Hedwig</h2>",
				"<p>Behold, the dwelling of God is with mankind. He will dwell with them, and they shall be his people, and God himself will be with them, and be their God.</p>",
				"<p>Be joyful in the Lord, all you lands; * <br/>",
				"<h3>OT Reading 1:1</h3>",
				"<h3>NT Reading 1:1</h3>",
				"<h3>Gospel Reading 1:1</h3>",
				"<em>Te Deum laudamus</em><br />",
			},
			wantStatus: http.StatusOK,
		},
		{
			name:       "favicon.ico",
			req:        httptest.NewRequest("GET", "http://testaddress/favicon.ico", nil),
			lectionary: &mockLectionary{},
			wantStatus: http.StatusOK,
		},

		{
			name: "error on keygen",
			keygen: func(time.Time) (calendar.KeyChain, error) {
				return calendar.KeyChain{}, errors.New("error on keygen")
			},
			req:        httptest.NewRequest("GET", "http://testaddress/?date=2014-12-26", nil),
			lectionary: &mockLectionary{},
			wantStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			templatePath := "/home/tmargheim/go/src/github.com/t-margheim/bcp-mp/cmd/mp.html"
			a := prayerApp{
				lectionaryService: tt.lectionary,
				page:              template.Must(template.ParseFiles(templatePath)),
				keyGenerator:      calendar.GetKeys,
			}
			a.ServeHTTP(w, tt.req)
			responseBody := w.Result().Body
			response, err := ioutil.ReadAll(responseBody)
			if err != nil {
				t.Errorf("failed to read response: %s", err)
			}
			responseString := string(response)

			if w.Code != tt.wantStatus {
				t.Errorf("unexpected response status code, got %d wanted %d", w.Code, tt.wantStatus)
			}

			for _, s := range tt.wantHTML {
				if !strings.Contains(responseString, s) {
					t.Errorf("missing expected html; response did not contain %s", s)
				}
			}
		})
	}
}

func Test_parseDate(t *testing.T) {
	tests := []struct {
		name     string
		r        *http.Request
		wantDate string
	}{
		{
			name:     "no query parameter provided",
			r:        httptest.NewRequest(http.MethodGet, "localhost:8080?", nil),
			wantDate: time.Now().Format("2006-01-02"),
		},
		{
			name:     "query parameter provided",
			r:        httptest.NewRequest(http.MethodGet, "localhost:8080?date=2019-02-28", nil),
			wantDate: "2019-02-28",
		},
		{
			name:     "invalid parameter provided",
			r:        httptest.NewRequest(http.MethodGet, "localhost:8080?date=badDate", nil),
			wantDate: time.Now().Format("2006-01-02"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseDate(tt.r)
			if got.Format("2006-01-02") != tt.wantDate {
				t.Errorf("parseDate() = %v, want %v", got, tt.wantDate)
			}
		})
	}
}
