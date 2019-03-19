// +build integration

package integration

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestIntegrationTest(t *testing.T) {
	tests := []struct {
		name             string
		date             string
		expectedElements []string
	}{
		{
			name: "December 26, 2018 - Success",
			date: "2018-12-26",
			expectedElements: []string{
				"<h1>Morning Prayer for December 26, 2018</h1>",
				"<h2>Saint Stephen, Deacon and Martyr</h2>",
				"<p>Behold, the dwelling of God is with mankind. He will dwell with them, and they shall be his people, and God himself will be with them, and be their God.</p>",
				"<p>Be joyful in the Lord, all you lands; * <br/>",
				"<h3>Psalm 28; Psalm 30</h3>",
				"<h3>2 Chronicles 24:17–22</h3>",
				"<h3>A Song of Praise</h3>",
				"<h3>Acts 6:1–7</h3>",
				"<h3>A Collect for Grace</h3>",
				"<em>Te Deum laudamus</em><br />",
				"<h3>For the Conservation of Natural Resources</h3>",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:8080?date=%s", tt.date))
			if err != nil {
				t.Errorf("error on request, is service running? error was: %s", err.Error())
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("error on reading response body: %s", err.Error())
			}

			bodyString := string(body)

			for _, s := range tt.expectedElements {
				if !strings.Contains(bodyString, s) {
					t.Errorf("missing expected html; response did not contain %s", s)
				}
			}
		})
	}
}
