package repository

import "go.mongodb.org/mongo-driver/mongo"

type ConfigurationRepository struct {
	ExercicioRepo *ExercicioRepository
}

func NewConfigurationRepositories(db *mongo.Database) *ConfigurationRepository {
	exercicioRepo := NewExercicioRepository(db)

	return &ConfigurationRepository{exercicioRepo}
}
