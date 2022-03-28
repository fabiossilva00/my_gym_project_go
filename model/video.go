package model

type Video struct {
	Link       string `bson:"link,omitempty"`
	Tipo_video string `bson:"tipo_video,omitempty"`
}
