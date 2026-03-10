package handler

import (
	"encoding/json"
	"net/http"

	"github.com/GrantMynott/gym/pkg/api/v1/exercise"
	"github.com/GrantMynott/gym/pkg/api/v1/workout"
)

func jsonHandler(body func(r *http.Request) any, statusCode int) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)

		jsonResponse, _ := json.Marshal(body(req))

		w.Write(jsonResponse)
	}
}

func Handler() {

	exerciseSvc := exercise.NewExerciseService()
	workoutApi := workout.NewSessionService()

	http.HandleFunc("GET /exercises", jsonHandler(func(r *http.Request) any {
		exercises := exerciseSvc.GetAll()
		return map[string]any{
			"exercises": exercises,
		}
	}, http.StatusOK))

	http.HandleFunc("POST /exercises", func(w http.ResponseWriter, r *http.Request) {
		var newExercise exercise.Exercise
		err := json.NewDecoder(r.Body).Decode(&newExercise)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		exerciseSvc.Add(newExercise)
		jsonHandler(func(req *http.Request) any {
			return newExercise
		}, http.StatusCreated)(w, r)
	})

	http.HandleFunc("GET /sessions", jsonHandler(func(r *http.Request) any {
		sessions := workoutApi.GetSessions()
		return map[string]any{
			"sessions": sessions,
		}
	}, http.StatusOK))

	http.ListenAndServe(":8080", nil)

}
