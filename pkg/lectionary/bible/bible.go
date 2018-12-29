package bible

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Service struct {
	BaseURL string
}

func (s *Service) GetLesson(reference string) Lesson {
	lessonString := url.QueryEscape(reference)
	// fmt.Println("lessonString:", lessonString)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(s.BaseURL, lessonString), nil)
	if err != nil {
		log.Println("failed to create request:", err.Error())
	}
	req.Header.Add("Authorization", "Token a9a234f364de585a1a6273b00ffe4be9c1b9ab47")
	httpResponse, _ := http.DefaultClient.Do(req)
	responseBody, _ := ioutil.ReadAll(httpResponse.Body)

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

	return Lesson{
		Reference: response.Canonical,
		Body:      template.HTML(body),
	}
}

type Lesson struct {
	Reference string
	Body      template.HTML
}
