package repository

import (
	"context"
	"fmt"
	"log"
	"my_gym_go/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

const nameDatabase = "exercicio"

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

	coll := e.database.Collection(nameDatabase)

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

func (e *ExercicioRepository) SalvarExercicio(exercicio model.ExercicioRequest) (model.Exercicio, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	exercicio.Created_at = time.Now()
	exercicio.Updated_at = time.Now()
	exercicioSave := model.Exercicio(exercicio)

	coll := e.database.Collection(nameDatabase)
	resultId, err := coll.InsertOne(ctx, exercicioSave)

	if err != nil {
		log.Println(err)
		return model.Exercicio{}, err
	}
	fmt.Printf("Criado novo exercicio com Id: %v", resultId.InsertedID)

	newId, _ := resultId.InsertedID.(primitive.ObjectID)

	exercicioSave.ID = newId

	return exercicioSave, nil
}
