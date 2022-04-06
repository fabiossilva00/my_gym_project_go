package service

import (
	"my_gym_go/model"
	"my_gym_go/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ExerciciosService struct {
	exercicioRepo *repository.ExercicioRepository
}

func NewExerciciosService(exercicioRepo *repository.ExercicioRepository) *ExerciciosService {
	return &ExerciciosService{exercicioRepo}
}

func (e *ExerciciosService) FindAll() []model.Exercicio {
	var exercicios []model.Exercicio

	exercicios = e.exercicioRepo.FindAll()

	return exercicios
}

func (e *ExerciciosService) SalvarExercicio(exercicio model.ExercicioRequest) (*model.Exercicio, error) {

	exercicioSalvo, err := e.exercicioRepo.SalvarExercicio(exercicio)

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return &exercicioSalvo, err
}
