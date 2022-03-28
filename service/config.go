package service

import "my_gym_go/repository"

type ConfigurationService struct {
	ExercicioService *ExerciciosService
}

func NewConfigurationService(r *repository.ConfigurationRepository) *ConfigurationService {

	exerciciosService := NewExerciciosService(r.ExercicioRepo)

	return &ConfigurationService{exerciciosService}
}
