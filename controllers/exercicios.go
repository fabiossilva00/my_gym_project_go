package controllers

import (
	"log"
	"my_gym_go/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ExerciciosController struct {
	exercicioService *service.ExerciciosService
}

func NewExerciciosController(er *service.ExerciciosService) *ExerciciosController {
	return &ExerciciosController{er}
}

func (e *ExerciciosController) FindAll(c echo.Context) error {
	exercicios := e.exercicioService.FindAll()
	log.Println(exercicios)

	return c.JSON(http.StatusOK, exercicios)

}
