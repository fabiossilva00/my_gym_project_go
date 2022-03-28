package service

import (
	"my_gym_go/model"
	"my_gym_go/repository"
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
