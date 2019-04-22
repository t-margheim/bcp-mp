package bible

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"google.golang.org/appengine/urlfetch"
)

type Service struct {
	BaseURL string
	client  *http.Client
}

func (s *Service) GetLesson(reference string) *Lesson {
	lessonString := url.QueryEscape(reference)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(s.BaseURL, lessonString), nil)
	if err != nil {
		return &Lesson{
			Reference: "Failed on http.NewRequest()",
			Body:      template.HTML(fmt.Sprintf("error message: %s", err.Error())),
		}
	}
	req.Header.Add("Authorization", "Token a9a234f364de585a1a6273b00ffe4be9c1b9ab47")
	httpResponse, err := s.client.Do(req)
	if err != nil {
		return &Lesson{
			Reference: "Failed on client.Do()",
			Body:      template.HTML(fmt.Sprintf("error message: %s", err.Error())),
		}
	}
	defer httpResponse.Body.Close()
	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return &Lesson{
			Reference: "Failed on ioutil.ReadAll()",
			Body:      template.HTML(fmt.Sprintf("error message: %s", err.Error())),
		}
	}

	var response resp
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		log.Println("unmarshal error:", err)
		fmt.Println(string(responseBody))
	}

	var body string
	for _, passage := range response.Passages {
		body += passage
	}

	return &Lesson{
		Reference: response.Canonical,
		Body:      template.HTML(body),
	}
}

func (s *Service) PrepareClient(ctx context.Context) {
	s.client = urlfetch.Client(ctx)
}

type Lesson struct {
	Reference string
	Body      template.HTML
}
