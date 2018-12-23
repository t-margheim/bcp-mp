package calendar

import "time"

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

	// OpenAdvent
	// OpenChristmas
	// OpenEpiphany
	// OpenLent
	// OpenHolyWeek
	// OpenEaster
	OpenTrinitySunday
	OpenAllSaints
	// OpenAnyTime
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
	}
)

type KeyChain struct {
	Season Key
	Open   Key
}

// GetKeys generates a KeyChain object for a given date. If the date is out of range,
// it will return an error
func GetKeys(date time.Time) (KeyChain, error) {
	var keys KeyChain
	keys.Season = GetSeason(date)
	keys.Open = GetOpen(date, keys.Season)

	return keys, nil
}

func GetSeason(date time.Time) Key {
	for _, dates := range seasons {
		if date.After(dates.start) && date.Before(dates.end) {
			return dates.season
		}
	}
	return Key(-1)
}

var specialOpens = map[string]Key{
	"2019-06-16": OpenTrinitySunday,
}

func GetOpen(date time.Time, season Key) Key {
	key, ok := specialOpens[date.Format("2006-01-02")]
	if ok {
		return key
	}

	return season
}
