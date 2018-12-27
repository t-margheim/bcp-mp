package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
	"github.com/t-margheim/bcp-mp/pkg/canticles"
	"github.com/t-margheim/bcp-mp/pkg/lectionary"

	"github.com/t-margheim/bcp-mp/pkg/opening"
)

func main() {
	app := prayerApp{}

	log.Fatal(http.ListenAndServe(":7777", &app))
}

type prayerApp struct {
}

const baseURL = "https://api.esv.org/v3/passage/html?include-verse-numbers=false&q=%s&include-footnotes=false&include-headings=false&include-first-verse-numbers=false&include-audio-link=false&include-chapter-numbers=false&include-passage-references=false&include-subheadings=false"

func (a *prayerApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	start := time.Now()

	date := time.Now().Add(-7 * time.Hour)
	keys, err := calendar.GetKeys(date)
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

	psalmReqStrings := []string{}
	for _, ps := range readings.Psalms {
		psalmReqStrings = append(psalmReqStrings, "Ps "+ps)
	}
	psalms := getLesson(strings.Join(psalmReqStrings, ";"))

	selectedDate := r.URL.Query().Get("date")
	if selectedDate != "" {
		newDate, err := time.Parse("2006-01-02", selectedDate)
		if err != nil {
			log.Println(err.Error())
		}
		date = newDate
	}

	cants := canticles.Get(keys)

	open, _ := opening.Get(date)
	elements := content{
		Date:       date.Format("2006-01-02"),
		Opening:    open,
		Invitatory: getInvitatory(keys),
		Psalms:     psalms,
		Canticle1:  cants[0],
		Canticle2:  cants[1],
		Gospel:     gospel,
		Lesson1:    firstLesson,
		Lesson2:    secondLesson,
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

func getInvitatory(keys calendar.KeyChain) invitatory {
	if keys.Season == calendar.SeasonEaster {
		return invitatory{
			Name: `Christ our Passover`,
			Content: `<p>Alleluia. <br/>
			Christ our Passover has been sacrificed for us; * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;therefore let us keep the feast, <br/>
			Not with old leaven, the leaven of malice and evil, * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;but with the unleavened bread of sincerity and truth. Alleluia. </p>
			
			<p>Christ being raised from the dead will never die again; * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;death no longer has dominion over him. <br/>
			The death that he died, he died to sin, once for all; * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;but the life he lives, he lives to God. <br/>
			So also consider yourselves dead to sin, * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;and alive to God in Jesus Christ our Lord. Alleluia.</p>
			
			<p>Christ has been raised from the dead, * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;the first fruits of those who have fallen asleep. <br/>
			For since by a man came death, * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;by a man has come also the resurrection of the dead. <br/>
			For as in Adam all die, * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;so in Christ shall all be made alive. Alleluia.</p>
			`,
		}
	}

	options := []invitatory{
		{
			Name: "Venite",
			Content: `<p>Come, let us sing to the Lord; * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;let us shout for joy to the Rock of our salvation. <br/>
			Let us come before his presence with thanksgiving * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;and raise a loud shout to him with psalms. </p>
			
			<p>For the Lord is a great God, * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;and a great King above all gods. <br/>
			In his hand are the caverns of the earth, * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;and the heights of the hills are his also. <br/>
			The sea is his, for he made it, * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;and his hands have molded the dry land.</p>
			
			<p>Come, let us bow down, and bend the knee, * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;and kneel before the Lord our Maker. <br/>
			For he is our God, <br/>
			and we are the people of his pasture and the sheep of his hand. *<br/>
			&nbsp;&nbsp;&nbsp;&nbsp;Oh, that today you would hearken to his voice!</p>
			`,
		},
		{
			Name: "Jubilate",
			Content: `<p>Be joyful in the Lord, all you lands; * <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;serve the Lord with gladness <br/>
			&nbsp;&nbsp;&nbsp;&nbsp;and come before his presence with a song. </p>
			<p>Know this: The Lord himself is God; * <br/>
&nbsp;&nbsp;&nbsp;&nbsp;he himself has made us, and we are his; <br/>
&nbsp;&nbsp;&nbsp;&nbsp;we are his people and the sheep of his pasture.</p>

<p>Enter his gates with thanksgiving; <br/>
go into his courts with praise; * <br/>
&nbsp;&nbsp;&nbsp;&nbsp;give thanks to him and call upon his Name. </p>

<p>For the Lord is good; <br/>
his mercy is everlasting; * <br/>
&nbsp;&nbsp;&nbsp;&nbsp;and his faithfulness endures from age to age. </p>
`,
		},
	}
	return options[keys.Iterator%2]
}

type lesson struct {
	Reference string
	Body      template.HTML
}

type content struct {
	Date       string
	Opening    opening.Opening
	Invitatory invitatory
	Psalms     lesson
	Canticle1  canticles.Canticle
	Canticle2  canticles.Canticle
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
