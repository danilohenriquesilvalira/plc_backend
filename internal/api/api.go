package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"plc_project/internal/database"

	"github.com/gorilla/mux"
)

// Usamos os tipos do pacote database como aliases.
type PLCMessage = database.PLC
type TagMessage = database.Tag

// getDBAndLogger cria a conexão com o banco e retorna o logger.
func getDBAndLogger() (*database.DB, *database.Logger, error) {
	db, err := database.NewDB(database.DBConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "apiuser",
		Password: "Danilo@34333528",
		Database: "plc_config",
	})
	if err != nil {
		return nil, nil, err
	}
	logger := database.NewLogger(db)
	return db, logger, nil
}

// GetPLCs retorna a lista de PLCs do banco.
func GetPLCs(w http.ResponseWriter, r *http.Request) {
	db, _, err := getDBAndLogger()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		log.Printf("Erro ao conectar ao banco: %v", err)
		return
	}
	defer db.Close()

	plcs, err := db.GetActivePLCs()
	if err != nil {
		http.Error(w, "Erro ao obter PLCs", http.StatusInternalServerError)
		log.Printf("Erro ao obter PLCs: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plcs)
}

// GetPLC retorna os detalhes de um PLC.
func GetPLC(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	db, _, err := getDBAndLogger()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	plcData, err := db.GetPLCByID(id)
	if err != nil {
		http.Error(w, "PLC não encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(plcData)
}

// CreatePLC cria um novo PLC e registra o evento.
func CreatePLC(w http.ResponseWriter, r *http.Request) {
	var newPLC PLCMessage
	if err := json.NewDecoder(r.Body).Decode(&newPLC); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	newPLC.LastUpdate = time.Now()
	db, logger, err := getDBAndLogger()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	createdPLC, err := db.CreatePLC(newPLC)
	if err != nil {
		http.Error(w, "Erro ao criar PLC", http.StatusInternalServerError)
		logger.Error("CreatePLC", fmt.Sprintf("Erro ao criar PLC: %v", err))
		return
	}
	logger.Info("CreatePLC", fmt.Sprintf("Novo PLC criado: %s (ID %d)", createdPLC.Name, createdPLC.ID))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdPLC)
}

// UpdatePLC atualiza um PLC existente e registra o evento.
func UpdatePLC(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	log.Printf("UpdatePLC: valor do parâmetro id: %q", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var plcUpdate PLCMessage
	if err := json.NewDecoder(r.Body).Decode(&plcUpdate); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	plcUpdate.ID = id
	plcUpdate.LastUpdate = time.Now()
	db, logger, err := getDBAndLogger()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := db.UpdatePLC(plcUpdate); err != nil {
		http.Error(w, "Erro ao atualizar PLC", http.StatusInternalServerError)
		logger.Error("UpdatePLC", fmt.Sprintf("Erro ao atualizar PLC: %v", err))
		return
	}
	logger.Info("UpdatePLC", fmt.Sprintf("PLC atualizado: ID %d, Nome %s", plcUpdate.ID, plcUpdate.Name))
	w.WriteHeader(http.StatusNoContent)
}

// DeletePLC remove um PLC e registra o evento.
func DeletePLC(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	db, logger, err := getDBAndLogger()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := db.DeletePLC(id); err != nil {
		http.Error(w, "Erro ao deletar PLC", http.StatusInternalServerError)
		logger.Error("DeletePLC", fmt.Sprintf("Erro ao deletar PLC: %v", err))
		return
	}
	logger.Info("DeletePLC", fmt.Sprintf("PLC deletado: ID %d", id))
	w.WriteHeader(http.StatusNoContent)
}

// GetTags retorna as tags de um PLC.
func GetTags(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	plcIDStr := vars["id"]
	plcID, err := strconv.Atoi(plcIDStr)
	if err != nil {
		http.Error(w, "PLC ID inválido", http.StatusBadRequest)
		return
	}
	db, _, err := getDBAndLogger()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	tags, err := db.GetPLCTags(plcID)
	if err != nil {
		http.Error(w, "Erro ao obter tags", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}

// CreateTag cria uma nova tag para um PLC e registra o evento.
func CreateTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	plcIDStr := vars["id"]
	plcID, err := strconv.Atoi(plcIDStr)
	if err != nil {
		http.Error(w, "PLC ID inválido", http.StatusBadRequest)
		return
	}
	var newTag TagMessage
	if err := json.NewDecoder(r.Body).Decode(&newTag); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	newTag.PLCID = plcID
	db, logger, err := getDBAndLogger()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	createdTag, err := db.CreateTag(newTag)
	if err != nil {
		http.Error(w, "Erro ao criar tag", http.StatusInternalServerError)
		logger.Error("CreateTag", fmt.Sprintf("Erro ao criar tag: %v", err))
		return
	}
	logger.Info("CreateTag", fmt.Sprintf("Tag criada: %s para PLC %d", createdTag.Name, createdTag.PLCID))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdTag)
}

// UpdateTag atualiza uma tag existente e registra o evento.
func UpdateTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tagIDStr := vars["tagId"]
	tagID, err := strconv.Atoi(tagIDStr)
	if err != nil {
		http.Error(w, "Tag ID inválido", http.StatusBadRequest)
		return
	}
	var tagUpdate TagMessage
	if err := json.NewDecoder(r.Body).Decode(&tagUpdate); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}
	tagUpdate.ID = tagID
	db, logger, err := getDBAndLogger()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := db.UpdateTag(tagUpdate); err != nil {
		http.Error(w, "Erro ao atualizar tag", http.StatusInternalServerError)
		logger.Error("UpdateTag", fmt.Sprintf("Erro ao atualizar tag: %v", err))
		return
	}
	logger.Info("UpdateTag", fmt.Sprintf("Tag atualizada: ID %d", tagUpdate.ID))
	w.WriteHeader(http.StatusNoContent)
}

// DeleteTag remove uma tag e registra o evento.
func DeleteTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tagIDStr := vars["tagId"]
	tagID, err := strconv.Atoi(tagIDStr)
	if err != nil {
		http.Error(w, "Tag ID inválido", http.StatusBadRequest)
		return
	}
	db, logger, err := getDBAndLogger()
	if err != nil {
		http.Error(w, "Erro ao conectar ao banco", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if err := db.DeleteTag(tagID); err != nil {
		http.Error(w, "Erro ao deletar tag", http.StatusInternalServerError)
		logger.Error("DeleteTag", fmt.Sprintf("Erro ao deletar tag: %v", err))
		return
	}
	logger.Info("DeleteTag", fmt.Sprintf("Tag deletada: ID %d", tagID))
	w.WriteHeader(http.StatusNoContent)
}
