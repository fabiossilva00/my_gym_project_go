package repository

import (
	"context"
	"log"
	"my_gym_go/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ExercicioRepository struct {
	database *mongo.Database
}

func NewExercicioRepository(database *mongo.Database) *ExercicioRepository {
	return &ExercicioRepository{database}
}

func (e *ExercicioRepository) FindAll() []model.Exercicio {
	var exercicios []model.Exercicio

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	coll := e.database.Collection("exercicio")

	filter := bson.D{{}}

	cur, err := coll.Find(ctx, filter)
	if err != nil {
		log.Panicln(err)
	}

	if err = cur.All(ctx, &exercicios); err != nil {
		log.Panicln(err)
	}

	defer cur.Close(ctx)

	return exercicios
}
