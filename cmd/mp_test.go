package main

import (
	"context"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
	"github.com/t-margheim/bcp-mp/pkg/lectionary"
	"github.com/t-margheim/bcp-mp/pkg/lectionary/bible"
)

func Test_prayerApp_ServeHTTP(t *testing.T) {
	tests := []struct {
		name       string
		req        *http.Request
		lectionary *lectionary.MockService
		wantHTML   []string
	}{
		{
			name: "Test 1",
			req:  httptest.NewRequest("GET", "http://testaddress/?date=2018-12-26", nil),
			lectionary: &lectionary.MockService{
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			templatePath := "/home/tmargheim/go/src/github.com/t-margheim/bcp-mp/morningprayer/mp.html"
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

			for _, s := range tt.wantHTML {
				if !strings.Contains(responseString, s) {
					t.Errorf("missing expected html; response did not contain %s", s)
				}
			}
		})
	}
}