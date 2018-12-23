package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/t-margheim/bcp-mp/pkg/opening"
)

func main() {
	// contents, err := ioutil.ReadFile("./gopher.json")
	// if err != nil {
	// 	log.Fatal("failed to read file:", err)
	// }
	// var story map[string]segment
	// err = json.Unmarshal(contents, &story)
	// if err != nil {
	// 	log.Fatal("failed to parse json:", err)
	// }

	app := prayerApp{}

	log.Fatal(http.ListenAndServe(":7777", &app))
}

type prayerApp struct {
}

func (a *prayerApp) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// // read path to determine which segment to show
	// path := r.URL.Path
	// chosen := "intro"

	// if len(path) > 1 {
	// 	chosen = path[1:]
	// }

	// segment, ok := a.content[chosen]
	// if !ok {
	// 	segment = a.content["intro"]
	// }

	date := time.Date(2018, 12, 17, 0, 0, 0, 0, time.UTC)
	open, _ := opening.Get(date)
	elements := content{
		Opening: open,
	}
	template := template.Must(template.ParseFiles("./mp.html"))

	template.Execute(w, elements)
	return
}

type content struct {
	Opening opening.Opening
}
