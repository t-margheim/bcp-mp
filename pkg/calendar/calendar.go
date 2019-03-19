package calendar

import (
	"errors"
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
	for _, dates := range seasons {
		if date.After(dates.start) && date.Before(dates.end) {
			d := date.Sub(dates.start)

			numWeeks := math.Floor(d.Hours() / 168)
			if dates.start.Weekday() == time.Sunday {
				numWeeks++
			}
			week := int(numWeeks)

			seasonEnd := dates.end
			if dates.season == SeasonAdvent {
				seasonEnd = seasonEnd.Add(192 * time.Hour)
			}

			year := seasonEnd.Year() % 2
			if year == 0 {
				year = 2
			}
			return KeyChain{
				Season: dates.season,
				Week:   week,
				Year:   year,
				Date:   date,
			}, nil
		}
	}
	return KeyChain{}, errors.New("date outside of calculated range")
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
