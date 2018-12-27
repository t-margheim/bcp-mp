package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_prayerApp_ServeHTTP(t *testing.T) {
	tests := []struct {
		name     string
		a        *prayerApp
		req      *http.Request
		wantHTML []string
	}{
		{
			name: "December 26",
			a:    &prayerApp{},
			req:  httptest.NewRequest("GET", "http://testaddress/?date=2018-12-26", nil),
			wantHTML: []string{
				"<h1>Morning Prayer for December 26, 2018</h1>",
				"<h2>Saint Stephen, Deacon and Martyr</h2>",
				"<p>Behold, I bring you good news of a great joy, which will come to all the people; for unto you is born this day in the city of David, a Savior, who is Christ the Lord.</p>",
				"<p>Be joyful in the Lord, all you lands; * <br/>",
				"<h3>2 Chronicles 24:17–22</h3>",
				"<em>Te Deum laudamus</em><br />",
			},
		},
		{
			name: "June 17",
			a:    &prayerApp{},
			req:  httptest.NewRequest("GET", "http://testaddress/?date=2019-06-17", nil),
			wantHTML: []string{
				"<h1>Morning Prayer for June 17, 2019</h1>",
				"<h2>Monday - The Season after Pentecost</h2>",
				"<p>Grace to you and peace from God our Father and from the Lord Jesus Christ.</p>",
				"<p>Come, let us sing to the Lord; * <br/>",
				"<h3>Ruth 1:1–18</h3>",
				"<em>Benedicite, omnia opera Domini</em><br />",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			tt.a.ServeHTTP(w, tt.req)
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
