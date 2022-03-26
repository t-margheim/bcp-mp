package opening

import (
	"github.com/t-margheim/bcp-mp/internal/calendar"
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

// Get returns a single opening verse that is valid based on the keychain passed in.
func Get(kc calendar.KeyChain) Opening {
	oo := Openings[kc.Season]
	if len(oo) == 0 {
		return Opening{}
	}

	return oo[kc.Iterator%len(oo)]
}
