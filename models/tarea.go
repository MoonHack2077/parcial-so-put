package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Tarea struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Titulo      string             `bson:"titulo" json:"titulo"`
	Descripcion string             `bson:"descripcion" json:"descripcion"`
	Completado  bool               `bson:"completado" json:"completado"`
	CreadoEn    time.Time          `bson:"creado_en" json:"creado_en"`
}
