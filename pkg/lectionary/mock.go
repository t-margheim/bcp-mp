package lectionary

import (
	"context"

	"github.com/t-margheim/bcp-mp/pkg/calendar"
)

type MockService struct {
	MockGetReadings func(context.Context, calendar.KeyChain) Readings
}

func (s *MockService) GetReadings(ctx context.Context, keys calendar.KeyChain) Readings {
	if s.MockGetReadings != nil {
		return s.MockGetReadings(ctx, keys)
	}
	return Readings{}
}
