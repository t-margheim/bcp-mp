package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
	"github.com/t-margheim/bcp-mp/pkg/lectionary"

	"github.com/t-margheim/bcp-mp/pkg/opening"
)

func main() {
	app := prayerApp{}

	log.Fatal(http.ListenAndServe(":7777", &app))
}

type prayerApp struct {
}

const baseURL = "https://api.esv.org/v3/passage/html?include-verse-numbers=false&q=%s&include-footnotes=false&include-headings=false&include-first-verse-number=false&include-audio-link=false&include-chapter-numbers=false&include-passage-references=false"

func (a *prayerApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	start := time.Now()

	keys, err := calendar.GetKeys(time.Now())
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	readings := lectionary.GetReadings(keys)
	fmt.Printf("%+v\n", readings)

	firstLesson := getLesson(readings.First)
	secondLesson := getLesson(readings.Second)
	gospel := getLesson(readings.Gospel)

	date := time.Now().Add(-7 * time.Hour)

	selectedDate := r.URL.Query().Get("date")
	if selectedDate != "" {
		newDate, err := time.Parse("2006-01-02", selectedDate)
		if err != nil {
			log.Println(err.Error())
		}
		date = newDate
	}

	open, _ := opening.Get(date)
	elements := content{
		Date:    date.Format("2006-01-02"),
		Opening: open,
		Invitatory: invitatory{
			Name: "Venite",
			Content: `Come, let us sing to the Lord; *<br>
			let us shout for joy to the Rock of our salvation. <br>
		Let us come before his presence with thanksgiving * <br>
			and raise a loud shout to him with psalms.<br>
		<br>
		For the Lord is a great God, * <br>
			and a great King above all gods. <br>
		In his hand are the caverns of the earth, * <br>
			and the heights of the hills are his also. <br>
		The sea is his, for he made it, * <br>
			and his hands have molded the dry land.<br>
		<br>
		Come, let us bow down, and bend the knee, * <br>
			and kneel before the Lord our Maker. <br>
		For he is our God, <br>
		and we are the people of his pasture and the sheep of his hand. *<br>
			Oh, that today you would hearken to his voice!`,
		},
		// Psalms:     psalms,
		Gospel:  gospel,
		Lesson1: firstLesson,
		Lesson2: secondLesson,
	}
	template := template.Must(template.ParseFiles("./mp.html"))

	template.Execute(w, elements)
	log.Println(r.URL.Path, "request served in", time.Since(start))
	return
}

func getLesson(reference string) lesson {
	lessonString := url.QueryEscape(reference)
	fmt.Println("lessonString:", lessonString)

	// lessonString = "Isa%2B42%3A1-12%3BEph%2B6%3A10-20%3BJohn%2B3%3A16-21"

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf(baseURL, lessonString), nil)
	req.Header.Add("Authorization", "Token a9a234f364de585a1a6273b00ffe4be9c1b9ab47")
	httpResponse, _ := http.DefaultClient.Do(req)
	responseBody, _ := ioutil.ReadAll(httpResponse.Body)

	var response resp
	err := json.Unmarshal(responseBody, &response)
	if err != nil {
		log.Println("unmarshal error:", err)
		fmt.Println(string(responseBody))
	}

	var body string
	for _, passage := range response.Passages {
		body += passage
	}

	return lesson{
		Reference: response.Canonical,
		Body:      template.HTML(body),
	}
}

type lesson struct {
	Reference string
	Body      template.HTML
}

type content struct {
	Date       string
	Opening    opening.Opening
	Invitatory invitatory
	Gospel     lesson
	Lesson1    lesson
	Lesson2    lesson
}

type invitatory struct {
	Name    string
	Content template.HTML
}

type resp struct {
	Canonical string   `json:"canonical"`
	Passages  []string `json:"passages"`
}
