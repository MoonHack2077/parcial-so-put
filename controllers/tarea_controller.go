package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MoonHack2077/parcial-so-put/models"
	"github.com/MoonHack2077/parcial-so-put/services"
	"github.com/gorilla/mux"
)

func CrearTarea(w http.ResponseWriter, r *http.Request) {
	var tarea models.Tarea
	if err := json.NewDecoder(r.Body).Decode(&tarea); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	id, err := services.CrearTarea(tarea)
	if err != nil {
		http.Error(w, "No se pudo crear la tarea", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id.Hex()})
}

func ObtenerTareas(w http.ResponseWriter, r *http.Request) {
	tareas, err := services.ObtenerTareas()
	if err != nil {
		http.Error(w, "Error al obtener las tareas", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tareas)
}

func ActualizarTarea(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var tarea models.Tarea
	if err := json.NewDecoder(r.Body).Decode(&tarea); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	err := services.ActualizarTarea(id, tarea)
	if err != nil {
		http.Error(w, "No se pudo actualizar la tarea", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Tarea actualizada correctamente"})
}
