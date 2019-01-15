package opening

import (
	"errors"
	"fmt"
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

// func init() {
// 	path := os.Getenv("OPENING_PATH")
// 	if path == "" {
// 		path = fmt.Sprintf("%s/src/github.com/t-margheim/bcp-mp/pkg/opening/data/", os.Getenv("GOPATH"))
// 	}
// 	for _, key := range keys {
// 		if _, ok := files[key]; !ok {
// 			continue
// 		}
// 		contents, err := ioutil.ReadFile(fmt.Sprintf("%s/%s.json", path, files[key]))
// 		if err != nil {
// 			log.Fatal("failed to read file", files[key], ":", err)
// 		}
// 		var openings []Opening
// 		err = json.Unmarshal(contents, &openings)
// 		if err != nil {
// 			log.Fatal("failed to parse json:", err)
// 		}
// 		Openings[key] = openings
// 	}
// }

// Opening defines the opening verse of the morning prayer.
type Opening struct {
	Text     string
	Citation string
}

var (
	files = map[calendar.Key]string{
		calendar.SeasonAdvent:      "advent",
		calendar.SeasonChristmas:   "christmas",
		calendar.SeasonEpiphany:    "epiphany",
		calendar.SeasonLent:        "lent",
		calendar.SeasonHolyWeek:    "holyweek",
		calendar.SeasonEaster:      "easter",
		calendar.SeasonOrdinary:    "ordinary",
		calendar.OpenTrinitySunday: "trinity",
		calendar.OpenAllSaints:     "saints",
	}
	// Openings = map[calendar.Key][]Opening{}
)

// Get returns a single opening verse that is valid based on the key passed in.
func Get(date time.Time) (Opening, error) {
	key := calendar.GetSeason(date).Season
	if key == calendar.Key(-1) {
		return Opening{}, fmt.Errorf("date %s is not in calendar lookup", date)
	}

	key = calendar.GetOpen(date, key)

	oo := Openings[key]
	if len(oo) == 0 {
		return Opening{}, errors.New("valid season, but no openings stored")
	}

	daysIterator := date.YearDay()

	return oo[daysIterator%len(oo)], nil
}
