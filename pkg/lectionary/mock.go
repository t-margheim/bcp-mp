package lectionary

import (
	"github.com/t-margheim/bcp-mp/pkg/calendar"
)

type MockService struct {
	MockGetReadings func(calendar.KeyChain) Readings
}

func (s *MockService) GetReadings(keys calendar.KeyChain) Readings {
	if s.MockGetReadings != nil {
		return s.MockGetReadings(keys)
	}
	return Readings{}
}
