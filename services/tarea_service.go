package services

import (
	"github.com/MoonHack2077/parcial-so-put/models"
	"github.com/MoonHack2077/parcial-so-put/repositories"
)

func ObtenerTareas() ([]models.Tarea, error) {
	return repositories.ObtenerTodasTareas()
}

func ActualizarTarea(id string, tarea models.Tarea) error {
	return repositories.ActualizarTarea(id, tarea)
}
