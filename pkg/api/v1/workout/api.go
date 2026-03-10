package workout

import (
	"github.com/GrantMynott/gym/pkg/api/v1/exercise"
)

type SessionService struct {
}

func NewSessionService() *SessionService {
	return &SessionService{}
}

func (w *SessionService) GetSessions() []Session {

	return []Session{
		{
			Name: "Morning Session",
			Workouts: []Workout{
				{
					Exercise:    exercise.Exercise{Name: "Push Up"},
					Repetitions: 10,
					Sets:        3,
				},
				{
					Exercise:    exercise.Exercise{Name: "Pull Up"},
					Repetitions: 6,
					Sets:        5,
				},
			},
		},
	}

}
