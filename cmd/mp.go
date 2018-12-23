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

	date := time.Now()

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
			Content: `Come, let us sing to the Lord; *
			let us shout for joy to the Rock of our salvation. 
		Let us come before his presence with thanksgiving * 
			and raise a loud shout to him with psalms.
		
		For the Lord is a great God, * 
			and a great King above all gods. 
		In his hand are the caverns of the earth, * 
			and the heights of the hills are his also. 
		The sea is his, for he made it, * 
			and his hands have molded the dry land.
		
		Come, let us bow down, and bend the knee, * 
			and kneel before the Lord our Maker. 
		For he is our God, 
		and we are the people of his pasture and the sheep of his hand. *
			Oh, that today you would hearken to his voice!`,
		},
	}
	template := template.Must(template.ParseFiles("./mp.html"))

	template.Execute(w, elements)
	return
}

type content struct {
	Date       string
	Opening    opening.Opening
	Invitatory invitatory
}

type invitatory struct {
	Name    string
	Content string
}
