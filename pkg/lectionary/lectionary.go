package lectionary

import (
	"context"
	"fmt"
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
	GetReadings(context.Context, calendar.KeyChain) Readings
}

type Service struct {
	bibleSvc bibleProvider
}

func New() *Service {
	return &Service{
		bibleSvc: &bible.Service{
			BaseURL: "https://api.esv.org/v3/passage/html?include-verse-numbers=false&q=%s&include-footnotes=false&include-headings=false&include-first-verse-numbers=false&include-audio-link=false&include-chapter-numbers=false&include-passage-references=false&include-subheadings=false",
		},
	}
}

func (s *Service) lookUpReferencesForDay(keys calendar.KeyChain) readingsReferences {
	var reading readingsReferences

	if special, ok := specialReadings[keys.ShortDate]; ok {
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

	for _, r := range dailyOffice[keys.Year] {
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

	return reading
}

func (s *Service) GetReadings(ctx context.Context, keys calendar.KeyChain) Readings {
	// figure out which passages are to be read that day
	passages := s.lookUpReferencesForDay(keys)

	// go get the text of each passage from ESV
	psalmReqStrings := []string{}
	for _, ps := range passages.Psalms {
		psalmReqStrings = append(psalmReqStrings, "Ps "+ps)
	}

	var first, second, gospel, psalms bible.Lesson

	s.bibleSvc.PrepareClient(ctx)

	finished := make(chan bool)
	go getLessonAsync(s.bibleSvc, passages.First, &first, finished)
	go getLessonAsync(s.bibleSvc, passages.Second, &second, finished)
	go getLessonAsync(s.bibleSvc, passages.Gospel, &gospel, finished)
	go getLessonAsync(s.bibleSvc, strings.Join(psalmReqStrings, ";"), &psalms, finished)

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

func getLessonAsync(bibleSvc bibleProvider, reference string, result *bible.Lesson, finished chan bool) {
	lesson := bibleSvc.GetLesson(reference)
	result.Body = lesson.Body
	result.Reference = lesson.Reference
	finished <- true
}

type bibleProvider interface {
	GetLesson(string) *bible.Lesson
	PrepareClient(context.Context)
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

type subReferences struct {
	First  string `json:"first"`
	Second string `json:"second"`
	Gospel string `json:"gospel"`
}
