package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercicio struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Nome           string             `bson:"nome"`
	Descricao      string             `bson:"descricao"`
	VariacoesNomes []string           `bson:"variacoes_nomes,omitempty"`
	Videos         []Video            `bson:"videos,omitempty"`
	Created_at     time.Time          `bson:"created_at"`
	Updated_at     time.Time          `bson:"updated_at"`
	Ativo          bool               `bson:"ativo"`
}
