package lectionary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
	"github.com/t-margheim/bcp-mp/pkg/lectionary/bible"
)

var (
	SeasonsLectionary = map[calendar.Key]string{
		calendar.SeasonAdvent:    "Advent",
		calendar.SeasonChristmas: "Christmas",
		calendar.SeasonEpiphany:  "Epiphany",
		calendar.SeasonLent:      "Lent",
		calendar.SeasonHolyWeek:  "Holy Week",
		calendar.SeasonEaster:    "Easter",
		calendar.SeasonOrdinary:  "The Season after Pentecost",
	}
)

type Provider interface {
	GetReadings(keys calendar.KeyChain) Readings
}

type Service struct {
	dailyOffice     map[int][]storedReadings
	specialReadings map[string]storedReadings
	// baseURL         string
	bibleSvc bible.Service
}

func New() *Service {

	svc := Service{
		dailyOffice:     map[int][]storedReadings{},
		specialReadings: map[string]storedReadings{},
		bibleSvc: bible.Service{
			BaseURL: "https://api.esv.org/v3/passage/html?include-verse-numbers=false&q=%s&include-footnotes=false&include-headings=false&include-first-verse-numbers=false&include-audio-link=false&include-chapter-numbers=false&include-passage-references=false&include-subheadings=false",
		},
	}

	// setup path to lectionary files
	path := os.Getenv("LECTIONARY_PATH")
	if path == "" {
		path = fmt.Sprintf("%s/src/github.com/t-margheim/bcp-mp/do-lect/daily-office/json/readings/", os.Getenv("GOPATH"))
	}
	for i := 1; i < 3; i++ {
		contents, err := ioutil.ReadFile(fmt.Sprintf("%s/dol-year-%d.min.json", path, i))
		if err != nil {
			log.Fatal("failed to read file", i, err)
		}
		var year []storedReadings
		err = json.Unmarshal(contents, &year)
		if err != nil {
			log.Fatal("failed to parse json:", err)
		}
		svc.dailyOffice[i] = year
	}
	contents, err := ioutil.ReadFile(fmt.Sprintf("%s/dol-holy-days.min.json", path))
	if err != nil {
		log.Fatal("failed to read holy days file", err)
	}
	var specials []storedReadings
	err = json.Unmarshal(contents, &specials)
	if err != nil {
		log.Fatal("failed to parse json:", err)
	}

	for _, ss := range specials {
		svc.specialReadings[ss.Day] = ss
	}

	return &svc
}

func (s *Service) lookUpReferencesForDay(keys calendar.KeyChain) readingsReferences {
	// fmt.Printf("%+v", keys)
	var reading readingsReferences

	if special, ok := s.specialReadings[keys.ShortDate]; ok {
		if special.Lessons.Morning != nil {
			special.Lessons = *special.Lessons.Morning
		}
		reading = readingsReferences{
			Psalms: special.Psalms.Morning,
			First:  special.Lessons.First,
			Second: special.Lessons.Second,
			Gospel: special.Lessons.Gospel,
			Title:  special.Title,
		}
		log.Println("special reading", reading)
		return reading
	}
	season := SeasonsLectionary[keys.Season]

	weekString := fmt.Sprintf("Week of %d %s", keys.Week, season)
	if keys.Season == calendar.SeasonOrdinary {
		weekString = fmt.Sprintf("Proper %d", keys.Week)
	}

	if keys.Season == calendar.SeasonChristmas {
		weekString = "Christmas Day and Following"
	}

	if keys.Season == calendar.SeasonEpiphany && keys.Week == 0 {
		weekString = "The Epiphany and Following"
	}

	for _, r := range s.dailyOffice[keys.Year] {
		if r.Season != season {
			continue
		}
		if r.Week != weekString {
			continue
		}

		if r.Day == keys.ShortDate {
			lessons := r.Lessons
			if lessons.Morning != nil {
				lessons = *r.Lessons.Morning
			}
			reading = readingsReferences{
				Psalms: r.Psalms.Morning,
				First:  lessons.First,
				Second: lessons.Second,
				Gospel: lessons.Gospel,
				Title:  r.Title,
			}
			log.Println("short date reading", reading)
			break
		}

		if r.Day == keys.Weekday {
			reading = readingsReferences{
				Psalms: r.Psalms.Morning,
				First:  r.Lessons.First,
				Second: r.Lessons.Second,
				Gospel: r.Lessons.Gospel,
			}
		}
	}

	log.Println("regular reading", reading)

	return reading
}

func (s *Service) GetReadings(keys calendar.KeyChain) Readings {
	// figure out which passages are to be read that day
	passages := s.lookUpReferencesForDay(keys)

	// go get the text of each passage from ESV
	psalmReqStrings := []string{}
	for _, ps := range passages.Psalms {
		psalmReqStrings = append(psalmReqStrings, "Ps "+ps)
	}

	var first, second, gospel, psalms bible.Lesson

	finished := make(chan bool)
	go s.getLessonAsync(passages.First, &first, finished)
	go s.getLessonAsync(passages.Second, &second, finished)
	go s.getLessonAsync(passages.Gospel, &gospel, finished)
	go s.getLessonAsync(strings.Join(psalmReqStrings, ";"), &psalms, finished)
	for i := 0; i < 4; i++ {
		<-finished
	}

	return Readings{
		First:  first,
		Second: second,
		Gospel: gospel,
		Psalms: psalms,
		Title:  passages.Title,
	}
}

func (s *Service) getLessonAsync(reference string, result *bible.Lesson, finished chan bool) {
	*result = s.bibleSvc.GetLesson(reference)
	finished <- true
}

type Readings struct {
	First, Second, Gospel, Psalms bible.Lesson
	Title                         string
}

type readingsReferences struct {
	Psalms                       []string
	First, Second, Gospel, Title string
}

type storedReadings struct {
	Year    string     `json:"year"`
	Season  string     `json:"season"`
	Week    string     `json:"week"`
	Day     string     `json:"day"`
	Title   string     `json:"title"`
	Psalms  psalm      `json:"psalms"`
	Lessons references `json:"lessons"`
}

type psalm struct {
	Morning []string `json:"morning"`
	Evening []string `json:"evening"`
}

type references struct {
	Morning *references `json:"morning"`
	Evening *references `json:"evening"`
	First   string      `json:"first"`
	Second  string      `json:"second"`
	Gospel  string      `json:"gospel"`
}
