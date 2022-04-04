package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Exercicio struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Nome           string             `bson:"nome" json:"nome" validate:"required,gte=3,lte=255"`
	Descricao      string             `bson:"descricao" json:"descricao,omitempty" validate:"required,gte=15,lte=255"`
	VariacoesNomes []string           `bson:"variacoes_nomes" json:"variacoes_nomes,omitempty"`
	Videos         []Video            `bson:"videos" json:"videos,omitempty"`
	Created_at     time.Time          `bson:"created_at" json:"created_at,omitempty" `
	Updated_at     time.Time          `bson:"updated_at" json:"updated_at,omitempty" `
	Ativo          bool               `bson:"ativo" json:"ativo,omitempty" validate:"required"`
}
