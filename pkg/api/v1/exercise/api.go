package exercise

import "sync"

type ExerciseService struct {
	mu        sync.Mutex
	exercises []Exercise
}

func NewExerciseService() *ExerciseService {
	return &ExerciseService{
		exercises: []Exercise{{Name: "Push Up"}, {Name: "Squat"}},
	}
}

func (s *ExerciseService) GetAll() []Exercise {
	s.mu.Lock()
	defer s.mu.Unlock()
	copySlice := make([]Exercise, len(s.exercises))
	copy(copySlice, s.exercises)
	return copySlice
}

// Add appends a new exercise to the list.
func (s *ExerciseService) Add(ex Exercise) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.exercises = append(s.exercises, ex)
}
