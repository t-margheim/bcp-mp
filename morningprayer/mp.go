package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/appengine"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
	"github.com/t-margheim/bcp-mp/pkg/canticles"
	"github.com/t-margheim/bcp-mp/pkg/lectionary"
	"github.com/t-margheim/bcp-mp/pkg/lectionary/bible"
	"github.com/t-margheim/bcp-mp/pkg/opening"
	"github.com/t-margheim/bcp-mp/pkg/prayers"
)

func main() {
	// set HTML template path
	templatePath := os.Getenv("TEMPLATE_PATH")
	if templatePath == "" {
		templatePath = "/home/tmargheim/go/src/github.com/t-margheim/bcp-mp/morningprayer/mp.html"
	}

	// set port configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	app := prayerApp{
		lectionaryService: lectionary.New(),
		page:              template.Must(template.ParseFiles(templatePath)),
	}

	log.Println("service is now running")
	http.Handle("/", &app)
	appengine.Main()
	// log.Fatal(http.ListenAndServe(port, &app))
}

type prayerApp struct {
	lectionaryService lectionary.Provider
	page              *template.Template
}

func (a *prayerApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	start := time.Now()

	date := time.Now().Add(-7 * time.Hour)
	selectedDate := r.URL.Query().Get("date")
	if selectedDate != "" {
		newDate, err := time.Parse("2006-01-02", selectedDate)
		if err != nil {
			log.Println(err.Error())
		}
		date = newDate
	}

	keys, err := calendar.GetKeys(date)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	title := fmt.Sprintf("%s - %s", keys.Weekday, lectionary.SeasonsLectionary[keys.Season])
	readings := a.lectionaryService.GetReadings(keys)

	if readings.Title != "" {
		title = readings.Title
	}

	cants := canticles.Get(keys)

	open, _ := opening.Get(date)
	elements := content{
		Date:       date.Format("January 2, 2006"),
		Title:      title,
		Opening:    open,
		Invitatory: getInvitatory(keys),
		Psalms:     readings.Psalms,
		Canticle1:  cants[0],
		Canticle2:  cants[1],
		Gospel:     readings.Gospel,
		Lesson1:    readings.First,
		Lesson2:    readings.Second,
		Suffrage:   prayers.Suffrages[keys.Iterator%len(prayers.Suffrages)],
		Collect:    prayers.Collects[keys.Weekday],
		Mission:    prayers.MissionPrayers[keys.Iterator%len(prayers.MissionPrayers)],
		Prayers:    prayers.DailyPrayers[keys.Iterator%len(prayers.DailyPrayers)],
		Closing:    closings[keys.Iterator%len(closings)],
	}

	a.page.Execute(w, elements)
	log.Println(r.URL.Path, "request served in", time.Since(start))
	return
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

type content struct {
	Date       string
	Title      string
	Opening    opening.Opening
	Invitatory invitatory
	Psalms     bible.Lesson
	Canticle1  canticles.Canticle
	Canticle2  canticles.Canticle
	Gospel     bible.Lesson
	Lesson1    bible.Lesson
	Lesson2    bible.Lesson
	Suffrage   prayers.Prayer
	Collect    prayers.Prayer
	Mission    prayers.Prayer
	Prayers    []prayers.Prayer
	Closing    string
}

type invitatory struct {
	Name    string
	Content template.HTML
}

var closings = []string{
	"The grace of our Lord Jesus Christ, and the love of God, and the fellowship of the Holy Spirit, be with us all evermore.",
	"May the God of hope fill us with all joy and peace in believing through the power of the Holy Spirit.",
	"Glory to God whose power, working in us, can do infinitely more than we can ask or imagine: Glory to him from generation to generation in the Church, and in Christ Jesus for ever and ever. ",
}
