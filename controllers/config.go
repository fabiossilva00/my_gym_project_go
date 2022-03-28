package controllers

import (
	"my_gym_go/service"
)

type ConfigurationControllers struct {
	ExercicioController *ExerciciosController
}

func NewConfigurationControllers(service *service.ConfigurationService) *ConfigurationControllers {

	exerciciosController := NewExerciciosController(service.ExercicioService)

	return &ConfigurationControllers{exerciciosController}
}
