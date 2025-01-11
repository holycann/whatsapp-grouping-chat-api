package folder

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/holycann/whatsapp-grouping-chat-api/models"
	"github.com/holycann/whatsapp-grouping-chat-api/utils"
)

type Handler struct {
	store models.FolderStore
}

func NewHandler(store models.FolderStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) FolderRoutes(router *mux.Router) {
	router.HandleFunc("/folder", h.HandleCreate).Methods("POST")
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var payload models.CreateFolderPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", errors))
		return
	}

	_, err := h.store.GetFolderByName(payload.Name)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("foldername %s already exists", payload.Name))
		return
	}

	err = h.store.CreateFolder(&models.CreateFolderPayload{
		ChatID:    payload.ChatID,
		Name:      payload.Name,
		CreatedAt: payload.CreatedAt,
	})
	if err != nil {
		fmt.Printf("error create folder: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var payload models.UpdateFolderPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		fmt.Printf("error parsing json: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	payload.ID = id

	if err := utils.Validate.Struct(payload); err != nil {
		fmt.Printf("error validating payload: %v\n", err)
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", errors))
		return
	}

	u, err := h.store.GetFolderByID(payload.ID)
	if err != nil {
		fmt.Printf("error get folder by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("folder id %d not found"))
		return
	}

	if u == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Folder with ID %d does not exist", payload.ID))
		return
	}

	if payload.ChatID <= 0 && payload.Name == "" {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Invalid ChatID And Phone Number Cannot Be Empty!"))
		return
	}

	err = h.store.UpdateFolder(&models.UpdateFolderPayload{
		ID:        payload.ID,
		ChatID:    payload.ChatID,
		Name:      payload.Name,
		UpdatedAt: payload.UpdatedAt,
	})
	if err != nil {
		fmt.Printf("error update chat: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Update folder %s successfully", u.Name))
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get folder by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("folder id %d not found"))
		return
	}

	err = h.store.DeleteFolder(id)
	if err != nil {
		fmt.Printf("error delete folder: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Delete folder successfully"))
}
