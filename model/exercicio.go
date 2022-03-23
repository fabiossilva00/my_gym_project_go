package model

import "time"

type Exercicio struct {
	_id            string    `bson:"title"`
	nome           string    `bson:"nome"`
	variacoesNomes []string  `bson:"variacoes_nomes"`
	created_at     time.Time `bson:"created_at"`
}
