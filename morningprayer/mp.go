package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
	"github.com/t-margheim/bcp-mp/pkg/canticles"
	"github.com/t-margheim/bcp-mp/pkg/invitatory"
	"github.com/t-margheim/bcp-mp/pkg/lectionary"
	"github.com/t-margheim/bcp-mp/pkg/lectionary/bible"
	"github.com/t-margheim/bcp-mp/pkg/opening"
	"github.com/t-margheim/bcp-mp/pkg/prayers"
	"google.golang.org/appengine"
)

func main() {
	// set HTML template path
	templatePath := os.Getenv("TEMPLATE_PATH")
	if templatePath == "" {
		templatePath = "./mp.html"
	}

	app := prayerApp{
		lectionaryService: lectionary.New(),
		page:              template.Must(template.ParseFiles(templatePath)),
		keyGenerator:      calendar.GetKeys,
	}

	log.Println("service is now running")
	http.Handle("/", &app)
	appengine.Main()
}

type prayerApp struct {
	lectionaryService lectionary.Provider
	page              *template.Template
	keyGenerator      func(time.Time) (calendar.KeyChain, error)
}

func (a *prayerApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
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

	keys, err := a.keyGenerator(date)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	title := fmt.Sprintf("%s - %s", keys.Weekday, lectionary.SeasonsLectionary[keys.Season])
	readings := a.lectionaryService.GetReadings(ctx, keys)

	if readings.Title != "" {
		title = readings.Title
	}

	cants := canticles.Get(keys)

	open, _ := opening.Get(date)
	elements := content{
		Date:       date.Format("January 2, 2006"),
		Title:      title,
		Opening:    open,
		Invitatory: invitatory.Get(keys),
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

type content struct {
	Date       string
	Title      string
	Opening    opening.Opening
	Invitatory invitatory.Entry
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

var closings = []string{
	"The grace of our Lord Jesus Christ, and the love of God, and the fellowship of the Holy Spirit, be with us all evermore.",
	"May the God of hope fill us with all joy and peace in believing through the power of the Holy Spirit.",
	"Glory to God whose power, working in us, can do infinitely more than we can ask or imagine: Glory to him from generation to generation in the Church, and in Christ Jesus for ever and ever. ",
}
