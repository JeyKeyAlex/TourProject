package saga

import (
	"github.com/rs/zerolog"

	pkgErr "github.com/JeyKeyAlex/TourProject/pkg/errors"
)

type ISaga interface {
	AddRollbackFunc(fn func() error) *Saga
	Rollback()
}

type Step struct {
	RollbackFunc func() error
}

type Saga struct {
	logger *zerolog.Logger
	steps  []Step
}

func (s *Saga) AddRollbackFunc(fn func() error) *Saga {
	s.steps = append(s.steps, Step{RollbackFunc: fn})
	return s
}

func (s *Saga) Rollback() {
	for i := len(s.steps) - 1; i >= 0; i-- {
		rollbackErr := s.steps[i].RollbackFunc()
		if rollbackErr != nil {
			s.logger.Error().Err(rollbackErr).Msg(pkgErr.ErrRollbackFailed)
		}
	}
}

func New() ISaga {
	return &Saga{
		steps: make([]Step, 0),
	}
}
