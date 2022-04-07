package controllers

import (
	"log"
	"my_gym_go/model"
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

func (e *ExerciciosController) SalvarExercicio(c echo.Context) error {
	exercicio := new(model.ExercicioRequest)

	if err := c.Bind(exercicio); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(exercicio); err != nil {
		log.Println(err)
		return err
	}

	var errorResponse model.Response[any]

	exercicioNew, err := e.exercicioService.SalvarExercicio(*exercicio)
	if err != nil {
		errorResponse.Data = err.Error()
		return echo.NewHTTPError(http.StatusInternalServerError, errorResponse)
	}

	return c.JSON(http.StatusCreated, exercicioNew)
}
