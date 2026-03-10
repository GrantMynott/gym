package workout

type Session struct {
	Name     string    `json:"name"`
	Workouts []Workout `json:"workouts"`
}
