package lectionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
)

var (
	doLectionary      = map[int][]storedReadings{}
	seasonsLectionary = map[calendar.Key]string{
		calendar.SeasonAdvent:    "Advent",
		calendar.SeasonChristmas: "Christmas",
		calendar.SeasonEpiphany:  "Epiphany",
		calendar.SeasonLent:      "Lent",
		calendar.SeasonHolyWeek:  "Holy Week",
		calendar.SeasonEaster:    "Easter",
		calendar.SeasonOrdinary:  "The Season after Pentecost",
	}
)

func init() {
	for i := 1; i < 3; i++ {
		contents, err := ioutil.ReadFile(fmt.Sprintf("%s/src/github.com/t-margheim/bcp-mp/do-lect/daily-office/json/readings/dol-year-%d.min.json", os.Getenv("GOPATH"), i))
		if err != nil {
			log.Fatal("failed to read file", i, err)
		}
		var year []storedReadings
		err = json.Unmarshal(contents, &year)
		if err != nil {
			log.Fatal("failed to parse json:", err)
		}
		doLectionary[i] = year
	}

}

func GetReadings(keys calendar.KeyChain) Readings {
	var reading Readings
	for _, r := range doLectionary[keys.Year] {
		season := seasonsLectionary[keys.Season]
		if r.Season != season {
			continue
		}
		weekString := fmt.Sprintf("Week of %d %s", keys.Week, season)
		if keys.Season == calendar.SeasonOrdinary {
			weekString = fmt.Sprintf("Proper %d", keys.Week)
		}
		if r.Week != weekString {
			continue
		}

		if r.Day == keys.ShortDate {
			lessons := r.Lessons
			if lessons.Morning != nil {
				lessons = *r.Lessons.Morning
			}
			reading = Readings{
				Psalms: r.Psalms.Morning,
				First:  lessons.First,
				Second: lessons.Second,
				Gospel: lessons.Gospel,
			}
			break
		}

		if r.Day == keys.Weekday {
			reading = Readings{
				Psalms: r.Psalms.Morning,
				First:  r.Lessons.First,
				Second: r.Lessons.Second,
				Gospel: r.Lessons.Gospel,
			}
		}
	}

	return reading
}

type Readings struct {
	Psalms                []string
	First, Second, Gospel string
}

type storedReadings struct {
	Year    string `json:"year"`
	Season  string `json:"season"`
	Week    string `json:"week"`
	Day     string `json:"day"`
	Title   string `json:"title"`
	Psalms  psalm  `json:"psalms"`
	Lessons lesson `json:"lessons"`
}

type psalm struct {
	Morning []string `json:"morning"`
	Evening []string `json:"evening"`
}

type lesson struct {
	Morning *lesson `json:"morning"`
	Evening *lesson `json:"evening"`
	First   string  `json:"first"`
	Second  string  `json:"second"`
	Gospel  string  `json:"gospel"`
}
