package lectionary

import (
	"context"
	"html/template"
	"log"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/t-margheim/bcp-mp/internal/calendar"
	"github.com/t-margheim/bcp-mp/pkg/lectionary/bible"
	"go.uber.org/zap"
)

func TestLookUpReferencesForDay(t *testing.T) {
	tests := []struct {
		name string
		keys calendar.KeyChain
		want readingsReferences
	}{
		{
			name: "December 17",
			keys: calendar.KeyChain{
				Season:    calendar.SeasonAdvent,
				Week:      3,
				Weekday:   "Monday",
				ShortDate: "Dec 17",
				Year:      1,
			},
			want: readingsReferences{
				Psalms: []string{"41", "52"},
				First:  "Isa 8:16–9:1",
				Second: "2 Pet 1:1–11",
				Gospel: "Luke 22:39–53",
			},
		},
		{
			name: "December 24",
			keys: calendar.KeyChain{
				Season:    calendar.SeasonAdvent,
				Week:      4,
				Weekday:   "Monday",
				ShortDate: "Dec 24",
				Year:      1,
			},
			want: readingsReferences{
				Psalms: []string{"45", "46"},
				First:  "Isa 35:1–10",
				Second: "Rev 22:12–17, 21",
				Gospel: "Luke 1:67–80",
				Title:  "Christmas Eve",
			},
		},
		{
			name: "December 26",
			keys: calendar.KeyChain{
				Season:    calendar.SeasonChristmas,
				Week:      0,
				Weekday:   "Wednesday",
				ShortDate: "Dec 26",
				Year:      1,
			},
			want: readingsReferences{
				Psalms: []string{"28", "30"},
				First:  "2 Chr 24:17–22",
				Second: "Acts 6:1–7",
				Gospel: "",
				Title:  "Saint Stephen, Deacon and Martyr",
			},
		},
		{
			name: "December 29",
			keys: calendar.KeyChain{
				Season:    calendar.SeasonChristmas,
				Week:      0,
				Weekday:   "Saturday",
				ShortDate: "Dec 29",
				Year:      1,
			},
			want: readingsReferences{
				Psalms: []string{"18:1–20"},
				First:  "Isa 12:1–6",
				Second: "Rev 1:1–8",
				Gospel: "John 7:37–52",
			},
		},
		{
			name: "January 7",
			keys: calendar.KeyChain{
				Season:    calendar.SeasonEpiphany,
				Week:      0,
				Weekday:   "Monday",
				ShortDate: "Jan 7",
				Year:      1,
			},
			want: readingsReferences{
				Psalms: []string{"103"},
				First:  "Isa 52:3–6",
				Second: "Rev 2:1–7",
				Gospel: "John 2:1–11",
			},
		},
		{
			name: "August 13",
			keys: calendar.KeyChain{
				Season:    calendar.SeasonOrdinary,
				Week:      9,
				Weekday:   "Tuesday",
				ShortDate: "Aug 13",
				Year:      1,
			},
			want: readingsReferences{
				Psalms: []string{"5", "6"},
				First:  "1 Sam 15:24–35",
				Second: "Acts 9:32–43",
				Gospel: "Luke 23:56b–24:11",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, _ := zap.NewDevelopment()
			svc := New(l.Sugar())
			if got := svc.lookUpReferencesForDay(tt.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lookUpReferencesForDay() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestService_lookUpReferencesForDay(t *testing.T) {
	type fields struct {
		dailyOffice     map[int][]storedReadings
		specialReadings map[string]storedReadings
		bibleSvc        bible.Service
	}
	type args struct {
		keys calendar.KeyChain
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   readingsReferences
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				// dailyOffice:     tt.fields.dailyOffice,
				// specialReadings: tt.fields.specialReadings,
				// bibleSvc: tt.fields.bibleSvc,
			}
			if got := s.lookUpReferencesForDay(tt.args.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.lookUpReferencesForDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLessonAsync(t *testing.T) {
	baseLesson := bible.Lesson{
		Reference: "Base Lesson",
		Body:      `<p></p>`,
	}

	tests := []struct {
		name           string
		bibleSvc       *mockBibleService
		reference      string
		expectedLesson bible.Lesson
	}{
		{
			name: "three hundred ms delay",
			bibleSvc: &mockBibleService{
				mockGetLesson: func(string) *bible.Lesson {
					time.Sleep(300 * time.Millisecond)
					return &baseLesson
				},
			},
			reference:      "testRef",
			expectedLesson: baseLesson,
		},
		{
			name: "three ms delay",
			bibleSvc: &mockBibleService{
				mockGetLesson: func(string) *bible.Lesson {
					time.Sleep(3 * time.Millisecond)
					return &baseLesson
				},
			},
			reference:      "other reference",
			expectedLesson: baseLesson,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			finished := make(chan bool)
			gotLesson := bible.Lesson{}
			go getLessonAsync(tt.bibleSvc, tt.reference, &gotLesson, finished)
			<-finished

			if !reflect.DeepEqual(gotLesson, tt.expectedLesson) {
				t.Errorf("bible service GetLesson did not modify passed in lesson, got %+v, wanted %+v", gotLesson, tt.expectedLesson)
			}

			if tt.bibleSvc.getLessonCalledTimes != 1 {
				t.Errorf("bible service GetLesson called wrong number of times, expected 1, got %d", tt.bibleSvc.getLessonCalledTimes)
			}

			if tt.bibleSvc.getLessonCalledWith[0] != tt.reference {
				t.Errorf("bible service GetLesson called with wrong value, expected %s, got %s", tt.reference, tt.bibleSvc.getLessonCalledWith)
			}
		})
	}
}

type mockBibleService struct {
	mockGetLesson        func(string) *bible.Lesson
	getLessonCalledTimes int
	getLessonCalledWith  []string

	mockPrepareClient        func(context.Context)
	prepareClientCalledTimes int
	prepareClientCalledWith  context.Context
}

func (s *mockBibleService) GetLesson(reference string) *bible.Lesson {
	s.getLessonCalledTimes++
	s.getLessonCalledWith = append(s.getLessonCalledWith, reference)

	if s.mockGetLesson != nil {
		return s.mockGetLesson(reference)
	}

	return nil
}

func (s *mockBibleService) PrepareClient(ctx context.Context) {
	s.prepareClientCalledTimes++
	s.prepareClientCalledWith = ctx

	if s.mockPrepareClient != nil {
		s.mockPrepareClient(ctx)
	}

}

func TestService_GetReadings(t *testing.T) {
	basePsalm := bible.Lesson{
		Reference: "Ps 72",
		Body:      template.HTML("psalm lesson"),
	}

	baseOT := bible.Lesson{
		Reference: "Jer 3:6–18",
		Body:      template.HTML("OT lesson"),
	}

	baseNT := bible.Lesson{
		Reference: "Rom 1:28–2:11",
		Body:      template.HTML("NT lesson"),
	}

	baseGospel := bible.Lesson{
		Reference: "John 5:1–18",
		Body:      template.HTML("gospel lesson"),
	}

	tests := []struct {
		name                    string
		bibleSvc                *mockBibleService
		keys                    calendar.KeyChain
		wantGetLessonCalledWith []string
		want                    Readings
	}{
		{
			name: "success - March 20, 2019",
			bibleSvc: &mockBibleService{
				mockPrepareClient: func(context.Context) {
					return
				},
				mockGetLesson: func(ref string) *bible.Lesson {
					switch ref {
					case basePsalm.Reference:
						return &basePsalm
					case baseOT.Reference:
						return &baseOT
					case baseNT.Reference:
						return &baseNT
					case baseGospel.Reference:
						return &baseGospel
					}
					return &bible.Lesson{}
				},
			},
			keys: calendar.KeyChain{
				Season:    4,
				Open:      4,
				Week:      2,
				Weekday:   "Wednesday",
				ShortDate: "Mar 20",
				Year:      1,
				Iterator:  443,
				Date:      time.Date(2019, 3, 20, 0, 0, 0, 0, time.UTC),
			},
			wantGetLessonCalledWith: []string{
				baseGospel.Reference,
				baseOT.Reference,
				basePsalm.Reference,
				baseNT.Reference,
			},
			want: Readings{
				First:  baseOT,
				Psalms: basePsalm,
				Second: baseNT,
				Gospel: baseGospel,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k, _ := calendar.GetKeys(time.Date(2019, 3, 20, 0, 0, 0, 0, time.UTC))
			log.Printf("%#v", k)
			ctx := context.Background()
			s := &Service{
				bibleSvc: tt.bibleSvc,
			}
			if got := s.GetReadings(ctx, tt.keys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetReadings() = %v, want %v", got, tt.want)
			}

			if tt.bibleSvc.prepareClientCalledTimes != 1 {
				t.Errorf("bibleSvc.prepareClient called wrong number of times, got %d, wanted 1", tt.bibleSvc.prepareClientCalledTimes)
			}

			if !reflect.DeepEqual(tt.bibleSvc.prepareClientCalledWith, ctx) {
				t.Errorf("bibleSvc.prepareClient called with wrong context, got %+v, wanted %+v", tt.bibleSvc.prepareClientCalledWith, ctx)
			}

			if tt.bibleSvc.getLessonCalledTimes != 4 {
				t.Errorf("bibleSvc.getLesson called wrong number of times, got %d, wanted 4", tt.bibleSvc.getLessonCalledTimes)
			}

			sort.Strings(tt.bibleSvc.getLessonCalledWith)
			sort.Strings(tt.wantGetLessonCalledWith)

			if !reflect.DeepEqual(tt.bibleSvc.getLessonCalledWith, tt.wantGetLessonCalledWith) {
				t.Errorf("bibleSvc.getLesson called with wrong references, got %+v, wanted %+v", tt.bibleSvc.getLessonCalledWith, tt.wantGetLessonCalledWith)
			}

		})
	}
}
