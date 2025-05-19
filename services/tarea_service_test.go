package services

import (
	// "os"
	"context"
	"testing"
	"time"

	// "github.com/joho/godotenv"
	"github.com/MoonHack2077/parcial-so-put/config"
	"github.com/MoonHack2077/parcial-so-put/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestObtenerTareas(t *testing.T) {
	tareas, err := ObtenerTareas()

	assert.NoError(t, err)
	assert.NotNil(t, tareas)
}

func TestActualizarTarea(t *testing.T) (primitive.ObjectID, error) {
	// Crear una tarea para actualizar
	tarea := models.Tarea{
		Titulo:      "Tarea para actualizar",
		Descripcion: "Antes de actualizar",
		Completado:  false,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tarea.CreadoEn = time.Now()
	collection := config.GetCollection("tareas")
	result, err := collection.InsertOne(ctx, tarea)

	assert.NoError(t, err)

	// Datos actualizados
	tareaActualizada := models.Tarea{
		Titulo:      "Tarea actualizada",
		Descripcion: "Después de actualizar",
		Completado:  true,
	}

	InsertedID := result.InsertedID.(primitive.ObjectID)
	err = ActualizarTarea(InsertedID.Hex(), tareaActualizada)
	assert.NoError(t, err)
	return InsertedID, err
}
