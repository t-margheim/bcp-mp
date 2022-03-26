package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/t-margheim/bcp-mp/internal/calendar"
	"github.com/t-margheim/bcp-mp/pkg/canticles"
	"github.com/t-margheim/bcp-mp/pkg/closings"
	"github.com/t-margheim/bcp-mp/pkg/invitatory"
	"github.com/t-margheim/bcp-mp/pkg/lectionary"
	"github.com/t-margheim/bcp-mp/pkg/lectionary/bible"
	"github.com/t-margheim/bcp-mp/pkg/opening"
	"github.com/t-margheim/bcp-mp/pkg/prayers"
	"go.uber.org/zap"
	"google.golang.org/appengine"
)

// go:embed mp.css
var css []byte

func main() {
	e := echo.New()
	zl, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	l := zl.Sugar()

	// set HTML template path
	templatePath := os.Getenv("TEMPLATE_PATH")
	if templatePath == "" {
		templatePath = "./mp.html"
	}

	app := prayerApp{
		keyGenerator:      calendar.GetKeys,
		l:                 l,
		lectionaryService: lectionary.New(l),
		page:              template.Must(template.ParseFiles(templatePath)),
	}

	l.Info("service is now running")
	// http.Handle("/", &app)
	// http.ListenAndServe(":80", &app)
	// appengine.Main()
	e.Renderer = &app
	e.Use(middleware.Logger())
	e.Static("static", "assets")
	e.GET("/", app.Handler)
	e.Start(":80")
}

type prayerApp struct {
	l                 *zap.SugaredLogger
	lectionaryService lectionary.Provider
	page              *template.Template
	keyGenerator      func(time.Time) (calendar.KeyChain, error)
}

// Render implements echo.Renderer
func (a *prayerApp) Render(w io.Writer, name string, elements interface{}, c echo.Context) error {

	return a.page.ExecuteTemplate(w, name, elements)
}

func (a *prayerApp) Handler(c echo.Context) error {
	date := parseDate(c.Request())

	keys, err := a.keyGenerator(date)
	if err != nil {
		return err
	}

	elements := a.generatePageContents(c.Request().Context(), keys)
	return c.Render(http.StatusOK, "mp.html", elements)
}

func (a *prayerApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "favicon.ico" {
		return
	}
	if r.URL.Path == "/mp.css" {
		a.l.Info("received css request")
		w.Write(css)
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

	a.l.Infow("request served",
		"path", r.URL.Path,
		"latency", time.Since(start),
		"readings", []string{
			elements.Psalms.Reference,
			elements.Lesson1.Reference,
			elements.Lesson2.Reference,
			elements.Gospel.Reference,
		},
	)
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
