package model

type Video struct {
	Link       string `bson:"link,omitempty" json:"link" validate:"required,url"`
	Tipo_video string `bson:"tipo_video,omitempty" json:"tipo_video" validate:"required"`
}
