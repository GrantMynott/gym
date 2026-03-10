package workout

import (
	"github.com/GrantMynott/gym/pkg/api/v1/exercise"
)

type Workout struct {
	Exercise    exercise.Exercise `json:"exercise"`
	Repetitions int8              `json:"repetitions"`
	Sets        int8              `json:"sets"`
}
