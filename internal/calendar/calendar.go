package calendar

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type Key int

const (
	_ Key = iota
	SeasonAdvent
	SeasonChristmas
	SeasonEpiphany
	SeasonLent
	SeasonHolyWeek
	SeasonEaster
	SeasonOrdinary

	OpenTrinitySunday
	OpenAllSaints
)

type dateRange struct {
	start, end time.Time
	season     Key
}

var (
	seasons = []dateRange{
		dateRange{
			start:  time.Date(2018, 12, 2, 0, 0, 0, 0, time.UTC),
			end:    time.Date(2018, 12, 25, 0, 0, 0, 0, time.UTC),
			season: SeasonAdvent,
		},

		dateRange{
			start:  time.Date(2018, 12, 25, 0, 0, 0, 0, time.UTC),
			end:    time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC),
			season: SeasonChristmas,
		},

		dateRange{
			start:  time.Date(2019, 1, 7, 0, 0, 0, 0, time.UTC),
			end:    time.Date(2019, 3, 6, 0, 0, 0, 0, time.UTC),
			season: SeasonEpiphany,
		},

		dateRange{
			start:  time.Date(2019, 3, 6, 0, 0, 0, 0, time.UTC),
			end:    time.Date(2019, 4, 14, 0, 0, 0, 0, time.UTC),
			season: SeasonLent,
		},

		dateRange{
			start:  time.Date(2019, 4, 14, 0, 0, 0, 0, time.UTC),
			end:    time.Date(2019, 4, 21, 0, 0, 0, 0, time.UTC),
			season: SeasonHolyWeek,
		},

		dateRange{
			start:  time.Date(2019, 4, 21, 0, 0, 0, 0, time.UTC),
			end:    time.Date(2019, 6, 9, 0, 0, 0, 0, time.UTC),
			season: SeasonEaster,
		},

		dateRange{
			start:  time.Date(2019, 6, 9, 0, 0, 0, 0, time.UTC),
			end:    time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC),
			season: SeasonOrdinary,
		},

		dateRange{
			start:  time.Date(2019, 12, 1, 0, 0, 0, 0, time.UTC),
			end:    time.Date(2019, 12, 25, 0, 0, 0, 0, time.UTC),
			season: SeasonAdvent,
		},
	}
)

type KeyChain struct {
	Season    Key
	Open      Key
	Week      int
	Weekday   string
	ShortDate string
	Year      int
	Iterator  int
	Date      time.Time
}

// GetKeys generates a KeyChain object for a given date. If the date is out of range,
// it will return an error
func GetKeys(date time.Time) (KeyChain, error) {
	keys, err := getSeason(date)
	if err != nil {
		return KeyChain{}, err
	}
	keys.Open = getOpen(date, keys.Season)
	keys.Weekday = date.Format("Monday")
	keys.ShortDate = date.Format("Jan 2")
	keys.Iterator = int(date.Sub(time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)).Hours() / 24)

	return keys, nil
}

func getSeason(date time.Time) (KeyChain, error) {
	yr := date.Year()
	easter, ok := easterLookup[yr]
	if !ok {
		return KeyChain{}, errors.New("date outside of calculated range")
	}

	if isLeapYear(yr) {
		easter.AddDate(0, 0, 1)
	}

	schedule, ok := yearlyScheduleLookupByEaster[easter.Format("January 2")]
	if !ok {
		return KeyChain{}, errors.New("could not retrieve schedule for year")
	}

	if isLeapYear(yr) {
		if schedule.AshWednesday.month == time.February {
			schedule.AshWednesday.day++
		}
	}

	fmt.Printf("%+v\n", schedule)

	var start time.Time
	var season Key
	switch {
	// still Christmas
	case date.Before(day(date.Year(), schedule.Epiphany.month, schedule.Epiphany.day)):
		start = day(date.Year()-1, time.December, 25)
		season = SeasonChristmas

	// epiphany
	case date.Before(day(date.Year(), schedule.AshWednesday.month, schedule.AshWednesday.day)):
		start = day(date.Year(), schedule.Epiphany.month, schedule.Epiphany.day)
		season = SeasonEpiphany

	// lent
	case date.Before(day(date.Year(), schedule.Easter.month, schedule.Easter.day)):
		start = day(date.Year(), schedule.AshWednesday.month, schedule.AshWednesday.day)
		season = SeasonLent

	// easter
	case date.Before(day(date.Year(), schedule.Pentecost.month, schedule.Pentecost.day)):
		start = day(date.Year(), schedule.Easter.month, schedule.Easter.day)
		season = SeasonEaster

	// pentecost
	case date.Before(day(date.Year(), schedule.Advent.month, schedule.Advent.day)):
		start = day(date.Year(), schedule.Pentecost.month, schedule.Pentecost.day)
		season = SeasonOrdinary

	// advent
	case date.After(day(date.Year(), schedule.Advent.month, schedule.Advent.day)):
		start = day(date.Year(), schedule.Advent.month, schedule.Advent.day)
		season = SeasonAdvent

	default:
		return KeyChain{}, fmt.Errorf("could not locate date on calendar, date %s, schedule: %+v", date.Format("Jan 2, 2006"), schedule)
	}

	d := date.Sub(start)

	numWeeks := math.Floor(d.Hours() / 168)
	if start.Weekday() == time.Sunday {
		numWeeks++
	}
	week := int(numWeeks)

	if season == SeasonAdvent || (season == SeasonChristmas && date.Month() == time.December) {
		yr++
	}
	return KeyChain{
		Season: season,
		Week:   week,
		Year:   yr % 2,
		Date:   date,
	}, nil
}

var specialOpens = map[string]Key{
	"2019-06-16": OpenTrinitySunday,
}

func getOpen(date time.Time, season Key) Key {
	key, ok := specialOpens[date.Format("2006-01-02")]
	if ok {
		return key
	}
	return season
}

func day(y int, m time.Month, d int) time.Time {
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

func isLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

var easterLookup = map[int]time.Time{
	2014: day(2014, time.April, 20),
	2015: day(2015, time.April, 5),
	2016: day(2016, time.March, 27),
	2017: day(2017, time.April, 16),
	2018: day(2018, time.April, 1),
	2019: day(2019, time.April, 21),
	2020: day(2020, time.April, 12),
	2021: day(2021, time.April, 4),
	2022: day(2022, time.April, 17),
	2023: day(2023, time.April, 9),
	2024: day(2024, time.March, 31),
	2025: day(2025, time.April, 20),
	2026: day(2026, time.April, 5),
	2027: day(2027, time.March, 28),
	2028: day(2028, time.April, 16),
	2029: day(2029, time.April, 1),
	2030: day(2030, time.April, 21),
	2031: day(2031, time.April, 13),
	2032: day(2032, time.March, 28),
	2033: day(2033, time.April, 17),
	2034: day(2034, time.April, 9),
	2035: day(2035, time.March, 25),
	2036: day(2036, time.April, 13),
	2037: day(2037, time.April, 5),
	2038: day(2038, time.April, 25),
	2039: day(2039, time.April, 10),
	2040: day(2040, time.April, 1),
	2041: day(2041, time.April, 21),
	2042: day(2042, time.April, 6),
	2043: day(2043, time.March, 29),
	2044: day(2044, time.April, 17),
	2045: day(2045, time.April, 9),
	2046: day(2046, time.March, 25),
	2047: day(2047, time.April, 14),
	2048: day(2048, time.April, 5),
	2049: day(2049, time.April, 18),
	2050: day(2050, time.April, 10),
	2051: day(2051, time.April, 2),
}

type monthDay struct {
	month time.Month
	day   int
}
type schedule struct {
	Epiphany             monthDay
	Easter               monthDay
	SundaysAfterEpiphany int
	AshWednesday         monthDay
	AscensionDay         monthDay
	Pentecost            monthDay
	ProperAfterPentecost int
	Advent               monthDay
}

var yearlyScheduleLookupByEaster = map[string]schedule{
	"March 22": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.March, 22},
		SundaysAfterEpiphany: 4,
		AshWednesday:         monthDay{time.February, 4},
		AscensionDay:         monthDay{time.April, 30},
		Pentecost:            monthDay{time.May, 10},
		ProperAfterPentecost: 3,
		Advent:               monthDay{time.November, 29},
	},
	"March 23": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.March, 23},
		SundaysAfterEpiphany: 4,
		AshWednesday:         monthDay{time.February, 5},
		AscensionDay:         monthDay{time.May, 1},
		Pentecost:            monthDay{time.May, 11},
		ProperAfterPentecost: 3,
		Advent:               monthDay{time.November, 30},
	},
	"March 24": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.March, 24},
		SundaysAfterEpiphany: 4,
		AshWednesday:         monthDay{time.February, 6},
		AscensionDay:         monthDay{time.May, 2},
		Pentecost:            monthDay{time.May, 12},
		ProperAfterPentecost: 3,
		Advent:               monthDay{time.December, 1},
	},
	"March 25": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.March, 25},
		SundaysAfterEpiphany: 5,
		AshWednesday:         monthDay{time.February, 7},
		AscensionDay:         monthDay{time.May, 3},
		Pentecost:            monthDay{time.May, 13},
		ProperAfterPentecost: 3,
		Advent:               monthDay{time.December, 2},
	},
	"March 26": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.March, 26},
		SundaysAfterEpiphany: 5,
		AshWednesday:         monthDay{time.February, 8},
		AscensionDay:         monthDay{time.May, 4},
		Pentecost:            monthDay{time.May, 14},
		ProperAfterPentecost: 3,
		Advent:               monthDay{time.December, 3},
	},
	"March 27": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.March, 27},
		SundaysAfterEpiphany: 5,
		AshWednesday:         monthDay{time.February, 9},
		AscensionDay:         monthDay{time.May, 5},
		Pentecost:            monthDay{time.May, 15},
		ProperAfterPentecost: 4,
		Advent:               monthDay{time.November, 27},
	},
	"March 28": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.March, 28},
		SundaysAfterEpiphany: 5,
		AshWednesday:         monthDay{time.February, 10},
		AscensionDay:         monthDay{time.May, 6},
		Pentecost:            monthDay{time.May, 16},
		ProperAfterPentecost: 4,
		Advent:               monthDay{time.November, 28},
	},
	"March 29": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.March, 29},
		SundaysAfterEpiphany: 5,
		AshWednesday:         monthDay{time.February, 11},
		AscensionDay:         monthDay{time.May, 7},
		Pentecost:            monthDay{time.May, 17},
		ProperAfterPentecost: 4,
		Advent:               monthDay{time.November, 29},
	},
	"March 30": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.March, 30},
		SundaysAfterEpiphany: 5,
		AshWednesday:         monthDay{time.February, 12},
		AscensionDay:         monthDay{time.May, 8},
		Pentecost:            monthDay{time.May, 18},
		ProperAfterPentecost: 4,
		Advent:               monthDay{time.November, 30},
	},
	"March 31": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.March, 31},
		SundaysAfterEpiphany: 5,
		AshWednesday:         monthDay{time.February, 13},
		AscensionDay:         monthDay{time.May, 9},
		Pentecost:            monthDay{time.May, 19},
		ProperAfterPentecost: 4,
		Advent:               monthDay{time.December, 1},
	},
	"April 1": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 1},
		SundaysAfterEpiphany: 6,
		AshWednesday:         monthDay{time.February, 14},
		AscensionDay:         monthDay{time.May, 10},
		Pentecost:            monthDay{time.May, 20},
		ProperAfterPentecost: 4,
		Advent:               monthDay{time.December, 2},
	},
	"April 2": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 2},
		SundaysAfterEpiphany: 6,
		AshWednesday:         monthDay{time.February, 15},
		AscensionDay:         monthDay{time.May, 11},
		Pentecost:            monthDay{time.May, 21},
		ProperAfterPentecost: 4,
		Advent:               monthDay{time.December, 3},
	},
	"April 3": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 3},
		SundaysAfterEpiphany: 6,
		AshWednesday:         monthDay{time.February, 16},
		AscensionDay:         monthDay{time.May, 12},
		Pentecost:            monthDay{time.May, 22},
		ProperAfterPentecost: 5,
		Advent:               monthDay{time.November, 27},
	},
	"April 4": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 4},
		SundaysAfterEpiphany: 6,
		AshWednesday:         monthDay{time.February, 17},
		AscensionDay:         monthDay{time.May, 13},
		Pentecost:            monthDay{time.May, 23},
		ProperAfterPentecost: 5,
		Advent:               monthDay{time.November, 28},
	},
	"April 5": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 5},
		SundaysAfterEpiphany: 6,
		AshWednesday:         monthDay{time.February, 18},
		AscensionDay:         monthDay{time.May, 14},
		Pentecost:            monthDay{time.May, 24},
		ProperAfterPentecost: 5,
		Advent:               monthDay{time.November, 29},
	},
	"April 6": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 6},
		SundaysAfterEpiphany: 6,
		AshWednesday:         monthDay{time.February, 19},
		AscensionDay:         monthDay{time.May, 15},
		Pentecost:            monthDay{time.May, 25},
		ProperAfterPentecost: 5,
		Advent:               monthDay{time.November, 30},
	},
	"April 7": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 7},
		SundaysAfterEpiphany: 6,
		AshWednesday:         monthDay{time.February, 20},
		AscensionDay:         monthDay{time.May, 16},
		Pentecost:            monthDay{time.May, 26},
		ProperAfterPentecost: 5,
		Advent:               monthDay{time.December, 1},
	},
	"April 8": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 8},
		SundaysAfterEpiphany: 7,
		AshWednesday:         monthDay{time.February, 21},
		AscensionDay:         monthDay{time.May, 17},
		Pentecost:            monthDay{time.May, 27},
		ProperAfterPentecost: 5,
		Advent:               monthDay{time.December, 2},
	},
	"April 9": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 9},
		SundaysAfterEpiphany: 7,
		AshWednesday:         monthDay{time.February, 22},
		AscensionDay:         monthDay{time.May, 18},
		Pentecost:            monthDay{time.May, 28},
		ProperAfterPentecost: 5,
		Advent:               monthDay{time.December, 3},
	},
	"April 10": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 10},
		SundaysAfterEpiphany: 7,
		AshWednesday:         monthDay{time.February, 23},
		AscensionDay:         monthDay{time.May, 19},
		Pentecost:            monthDay{time.May, 29},
		ProperAfterPentecost: 6,
		Advent:               monthDay{time.November, 27},
	},
	"April 11": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 11},
		SundaysAfterEpiphany: 7,
		AshWednesday:         monthDay{time.February, 24},
		AscensionDay:         monthDay{time.May, 20},
		Pentecost:            monthDay{time.May, 30},
		ProperAfterPentecost: 6,
		Advent:               monthDay{time.November, 28},
	},
	"April 12": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 12},
		SundaysAfterEpiphany: 7,
		AshWednesday:         monthDay{time.February, 25},
		AscensionDay:         monthDay{time.May, 21},
		Pentecost:            monthDay{time.May, 31},
		ProperAfterPentecost: 6,
		Advent:               monthDay{time.November, 29},
	},
	"April 13": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 13},
		SundaysAfterEpiphany: 7,
		AshWednesday:         monthDay{time.February, 26},
		AscensionDay:         monthDay{time.May, 22},
		Pentecost:            monthDay{time.June, 1},
		ProperAfterPentecost: 6,
		Advent:               monthDay{time.November, 30},
	},
	"April 14": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 14},
		SundaysAfterEpiphany: 7,
		AshWednesday:         monthDay{time.February, 27},
		AscensionDay:         monthDay{time.May, 23},
		Pentecost:            monthDay{time.June, 2},
		ProperAfterPentecost: 6,
		Advent:               monthDay{time.December, 1},
	},
	"April 15": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 15},
		SundaysAfterEpiphany: 8,
		AshWednesday:         monthDay{time.February, 28},
		AscensionDay:         monthDay{time.May, 24},
		Pentecost:            monthDay{time.June, 3},
		ProperAfterPentecost: 6,
		Advent:               monthDay{time.December, 2},
	},
	"April 16": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 16},
		SundaysAfterEpiphany: 8,
		AshWednesday:         monthDay{time.March, 1},
		AscensionDay:         monthDay{time.May, 25},
		Pentecost:            monthDay{time.June, 4},
		ProperAfterPentecost: 6,
		Advent:               monthDay{time.December, 3},
	},
	"April 17": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 17},
		SundaysAfterEpiphany: 8,
		AshWednesday:         monthDay{time.March, 2},
		AscensionDay:         monthDay{time.May, 26},
		Pentecost:            monthDay{time.June, 5},
		ProperAfterPentecost: 7,
		Advent:               monthDay{time.November, 27},
	},
	"April 18": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 18},
		SundaysAfterEpiphany: 8,
		AshWednesday:         monthDay{time.March, 3},
		AscensionDay:         monthDay{time.May, 27},
		Pentecost:            monthDay{time.June, 6},
		ProperAfterPentecost: 7,
		Advent:               monthDay{time.November, 28},
	},
	"April 19": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 19},
		SundaysAfterEpiphany: 8,
		AshWednesday:         monthDay{time.March, 4},
		AscensionDay:         monthDay{time.May, 28},
		Pentecost:            monthDay{time.June, 7},
		ProperAfterPentecost: 7,
		Advent:               monthDay{time.November, 29},
	},
	"April 20": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 20},
		SundaysAfterEpiphany: 8,
		AshWednesday:         monthDay{time.March, 5},
		AscensionDay:         monthDay{time.May, 29},
		Pentecost:            monthDay{time.June, 8},
		ProperAfterPentecost: 7,
		Advent:               monthDay{time.November, 30},
	},
	"April 21": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 21},
		SundaysAfterEpiphany: 8,
		AshWednesday:         monthDay{time.March, 6},
		AscensionDay:         monthDay{time.May, 30},
		Pentecost:            monthDay{time.June, 9},
		ProperAfterPentecost: 7,
		Advent:               monthDay{time.December, 1},
	},
	"April 22": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 22},
		SundaysAfterEpiphany: 9,
		AshWednesday:         monthDay{time.March, 7},
		AscensionDay:         monthDay{time.May, 31},
		Pentecost:            monthDay{time.June, 10},
		ProperAfterPentecost: 7,
		Advent:               monthDay{time.December, 2},
	},
	"April 23": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 23},
		SundaysAfterEpiphany: 9,
		AshWednesday:         monthDay{time.March, 8},
		AscensionDay:         monthDay{time.June, 1},
		Pentecost:            monthDay{time.June, 11},
		ProperAfterPentecost: 7,
		Advent:               monthDay{time.December, 3},
	},
	"April 24": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 24},
		SundaysAfterEpiphany: 9,
		AshWednesday:         monthDay{time.March, 9},
		AscensionDay:         monthDay{time.June, 2},
		Pentecost:            monthDay{time.June, 12},
		ProperAfterPentecost: 8,
		Advent:               monthDay{time.November, 27},
	},
	"April 25": {
		Epiphany:             monthDay{time.January, 7},
		Easter:               monthDay{time.April, 25},
		SundaysAfterEpiphany: 9,
		AshWednesday:         monthDay{time.March, 10},
		AscensionDay:         monthDay{time.June, 3},
		Pentecost:            monthDay{time.June, 13},
		ProperAfterPentecost: 8,
		Advent:               monthDay{time.November, 28},
	},
}

/*
		Sundays	 	 	 	Numbered
Easter	after	Ash			Ascension
Day		Epiph	Wednesday	Day			Pentecost			Advent
	"March 22": {March 22	4	Feb. 4		April 30	May 10		#3		November 29},
	"March 23": {March 23	4	Feb. 5		May 1		May 11		#3		November 30},
	"March 24": {March 24	4	Feb. 6		May 2		May 12		#3		December 1},
	"March 25": {March 25	5	Feb. 7		May 3		May 13		#3		December 2},
	"March 26": {March 26	5	Feb. 8		May 4		May 14		#3		December 3},
	"March 27": {March 27	5	Feb. 9		May 5		May 15		#4		November 27},
	"March 28": {March 28	5	Feb. 10		May 6		May 16		#4		November 28},
	"March 29": {March 29	5	Feb. 11		May 7		May 17		#4		November 29},
	"March 30": {March 30	5	Feb. 12		May 8		May 18		#4		November 30},
	"March 31": {March 31	5	Feb. 13		May 9		May 19		#4		December 1},
	"April 1	": {April 1		6	Feb. 14		May 10		May 20		#4		December 2},
	"April 2	": {April 2		6	Feb. 15		May 11		May 21		#4		December 3},
	"April 3	": {April 3		6	Feb. 16		May 12		May 22		#5		November 27},
	"April 4	": {April 4		6	Feb. 17		May 13		May 23		#5		November 28},
	"April 5	": {April 5		6	Feb. 18		May 14		May 24		#5		November 29},
	"April 6	": {April 6		6	Feb. 19		May 15		May 25		#5		November 30},
	"April 7	": {April 7		6	Feb. 20		May 16		May 26		#5		December 1},
	"April 8	": {April 8		7	Feb. 21		May 17		May 27		#5		December 2},
	"April 9	": {April 9		7	Feb. 22		May 18		May 28		#5		December 3},
	"April 10": {April 10	7	Feb. 23		May 19		May 29		#6		November 27},
	"April 11": {April 11	7	Feb. 24		May 20		May 30		#6		November 28},
	"April 12": {April 12	7	Feb. 25		May 21		May 31		#6		November 29},
	"April 13": {April 13	7	Feb. 26		May 22		June 1		#6		November 30},
	"April 14": {April 14	7	Feb. 27		May 23		June 2		#6		December 1},
	"April 15": {April 15	8	Feb. 28		May 24		June 3		#6		December 2},
	"April 16": {April 16	8	March 1		May 25		June 4		#6		December 3},
	"April 17": {April 17	8	March 2		May 26		June 5		#7		November 27},
	"April 18": {April 18	8	March 3		May 27		June 6		#7		November 28},
	"April 19": {April 19	8	March 4		May 28		June 7		#7		November 29},
	"April 20": {April 20	8	March 5		May 29		June 8		#7		November 30},
	"April 21": {April 21	8	March 6		May 30		June 9		#7		December 1},
	"April 22": {April 22	9	March 7		May 31		June 10		#7		December 2},
	"April 23": {April 23	9	March 8		June 1		June 11		#7		December 3},
	"April 24": {April 24	9	March 9		June 2		June 12		#8		November 27},
	"April 25": {April 25	9	March 10	June 3		June 13		#8		November 28},
*/
