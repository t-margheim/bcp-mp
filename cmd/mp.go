package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
	"github.com/t-margheim/bcp-mp/pkg/canticles"
	"github.com/t-margheim/bcp-mp/pkg/closings"
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

	if r.URL.Path == "/favicon.ico" {
		return
	}

	start := time.Now()

	ctx := appengine.NewContext(r)

	date := parseDate(r)

	keys, err := a.keyGenerator(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	elements := a.generatePageContents(ctx, keys)

	a.page.Execute(w, elements)

	log.Println(r.URL.Path, "request served in", time.Since(start))
	return
}

func parseDate(r *http.Request) time.Time {
	date := time.Now().UTC()
	selectedDate := r.URL.Query().Get("date")
	if selectedDate != "" {
		newDate, err := time.Parse("2006-01-02", selectedDate)
		if err != nil {
			log.Println(err.Error())
		} else {
			date = newDate
		}
	}
	return date
}

func (a *prayerApp) generatePageContents(ctx context.Context, keys calendar.KeyChain) content {
	title := fmt.Sprintf("%s - %s", keys.Weekday, lectionary.SeasonsLectionary[keys.Season])
	readings := a.lectionaryService.GetReadings(ctx, keys)

	if readings.Title != "" {
		title = readings.Title
	}

	cants := canticles.Get(keys.Iterator)

	return content{
		Date:       keys.Date.Format("January 2, 2006"),
		Title:      title,
		Opening:    opening.Get(keys),
		Invitatory: invitatory.Get(keys),
		Psalms:     readings.Psalms,
		Canticle1:  cants[0],
		Canticle2:  cants[1],
		Gospel:     readings.Gospel,
		Lesson1:    readings.First,
		Lesson2:    readings.Second,
		Suffrage:   prayers.GetSuffrage(keys.Iterator),
		Collect:    prayers.GetDailyCollect(keys.Weekday),
		Mission:    prayers.GetPrayerForMission(keys.Iterator),
		Prayers:    prayers.GetDailyPrayers(keys.Iterator),
		Closing:    closings.Get(keys.Iterator),
	}
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
