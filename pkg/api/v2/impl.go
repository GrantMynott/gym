package api

import (
	"context"
	"database/sql"
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

//go:embed index.html
//go:embed openapi.yaml

var swaggerUI embed.FS

type Server struct {
	db *sql.DB
}

func NewServer(e *echo.Echo, db *sql.DB) *Server {
	e.GET("/swagger/*", echo.WrapHandler(http.StripPrefix("/swagger/", http.FileServer(http.FS(swaggerUI)))))

	return &Server{
		db: db,
	}
}

func (s *Server) GetExercises(ctx context.Context, request GetExercisesRequestObject) (GetExercisesResponseObject, error) {

	return GetExercises200JSONResponse([]Exercise{}), nil
}

func (s *Server) GetExercise(ctx context.Context, request GetExerciseRequestObject) (GetExerciseResponseObject, error) {

	return GetExercise200JSONResponse(Exercise{}), nil
}

func (s *Server) PostExercise(ctx context.Context, request PostExerciseRequestObject) (PostExerciseResponseObject, error) {

	exercise_query := `INSERT INTO gym.exercise (name) VALUES ($1) RETURNING id`

	var id int
	err := s.db.QueryRow(exercise_query, request.Body.Name).Scan(&id)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	for _, v := range request.Body.Muscles {
		muscles_query := `INSERT INTO gym.exercise_muscles (exercise_id, muscle) VALUES ($1,$2)`
		s.db.Exec(muscles_query, id, v)
	}
	return PostExercise201JSONResponse(Exercise{
		Name:    request.Body.Name,
		Muscles: request.Body.Muscles,
	}), nil
}
