package opening

import (
	"errors"
	"fmt"
	"time"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
)

// Opening defines the opening verse of the morning prayer.
type Opening struct {
	Text     string
	Citation string
}

var (
	Openings = map[calendar.Key][]Opening{
		calendar.SeasonAdvent: []Opening{
			{
				Text:     "Watch, for you know not when the master of the house will come, in the evening, or at midnight, or at cockcrow, or in the morning; lest he come suddenly and find you asleep.",
				Citation: "Mark 13:35, 36",
			},
			{
				Text:     "In the wilderness prepare the way of the Lord, make straight in the desert a highway for our God.",
				Citation: "Isaiah 40:3",
			},
			{
				Text:     "The glory of the Lord shall be revealed, and all flesh shall see it together.",
				Citation: "Isaiah 40:5",
			},
		},

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
