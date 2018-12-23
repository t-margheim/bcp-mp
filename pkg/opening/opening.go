package opening

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
)

var (
	keys = []calendar.Key{
		calendar.SeasonAdvent,
		calendar.SeasonChristmas,
		calendar.SeasonEpiphany,
		calendar.SeasonLent,
		calendar.SeasonHolyWeek,
		calendar.SeasonEaster,
		calendar.SeasonOrdinary,
		calendar.OpenTrinitySunday,
		calendar.OpenAllSaints,
	}
)

func init() {
	for _, key := range keys {
		if _, ok := files[key]; !ok {
			continue
		}
		contents, err := ioutil.ReadFile(fmt.Sprintf("./data/%s.json", files[key]))
		if err != nil {
			log.Fatal("failed to read file", files[key], ":", err)
		}
		var openings []Opening
		err = json.Unmarshal(contents, &openings)
		if err != nil {
			log.Fatal("failed to parse json:", err)
		}
		Openings[key] = openings
	}
}

// Opening defines the opening verse of the morning prayer.
type Opening struct {
	Text     string
	Citation string
}

var (
	files = map[calendar.Key]string{
		calendar.SeasonAdvent: "advent",
		// calendar.SeasonChristmas:   "christmas",
		// calendar.SeasonEpiphany:    "epiphany",
		// calendar.SeasonLent:        "lent",
		// calendar.SeasonHolyWeek:    "holyweek",
		// calendar.SeasonEaster:      "easter",
		// calendar.SeasonOrdinary:    "ordinary",
		// calendar.OpenTrinitySunday: "trinity",
		// calendar.OpenAllSaints:     "saints",
	}
	Openings = map[calendar.Key][]Opening{
		calendar.OpenTrinitySunday: []Opening{
			{
				Text:     "Holy, holy, holy is the Lord God Almighty, who was, and is, and is to come.",
				Citation: "Revelation 4:8",
			},
		},
	}
)

// Get returns a single opening verse that is valid based on the key passed in.
func Get(date time.Time) (Opening, error) {
	key := calendar.GetSeason(date)
	if key == calendar.Key(-1) {
		return Opening{}, fmt.Errorf("date %s is not in calendar lookup", date)
	}

	fmt.Println(key)

	key = calendar.GetOpen(date, key)
	fmt.Println(key)

	oo := Openings[key]
	if len(oo) == 0 {
		return Opening{}, errors.New("valid season, but no openings stored")
	}

	daysIterator := date.YearDay()

	return oo[daysIterator%len(oo)], nil
}
