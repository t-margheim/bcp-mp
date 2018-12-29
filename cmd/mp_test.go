package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
	"github.com/t-margheim/bcp-mp/pkg/lectionary/bible"

	"github.com/t-margheim/bcp-mp/pkg/lectionary"
)

func Test_prayerApp_ServeHTTP(t *testing.T) {
	tests := []struct {
		name       string
		req        *http.Request
		lectionary lectionary.Provider
		wantHTML   []string
	}{
		{
			name: "Test 1",
			req:  httptest.NewRequest("GET", "http://testaddress/?date=2018-12-26", nil),
			lectionary: &lectionary.MockService{
				MockGetReadings: func(calendar.KeyChain) lectionary.Readings {
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
				"<p>Behold, I bring you good news of a great joy, which will come to all the people; for unto you is born this day in the city of David, a Savior, who is Christ the Lord.</p>",
				"<p>Be joyful in the Lord, all you lands; * <br/>",
				"<h3>OT Reading 1:1</h3>",
				"<h3>NT Reading 1:1</h3>",
				"<h3>Gospel Reading 1:1</h3>",
				"<em>Te Deum laudamus</em><br />",
			},
		},
		// {
		// 	name: "June 17",
		// 	req:  httptest.NewRequest("GET", "http://testaddress/?date=2019-06-17", nil),
		// 	lectionary: &lectionary.MockService{
		// 		MockGetReadings: func(calendar.KeyChain) lectionary.Readings {
		// 			return lectionary.Readings{}
		// 		},
		// 	},
		// 	wantHTML: []string{
		// 		"<h1>Morning Prayer for June 17, 2019</h1>",
		// 		"<h2>Monday - The Season after Pentecost</h2>",
		// 		"<p>Grace to you and peace from God our Father and from the Lord Jesus Christ.</p>",
		// 		"<p>Come, let us sing to the Lord; * <br/>",
		// 		"<h3>Ruth 1:1–18</h3>",
		// 		"<em>Benedicite, omnia opera Domini</em><br />",
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			a := prayerApp{
				lectionaryService: tt.lectionary,
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
