package chat

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
	store models.ChatStore
}

func NewHandler(store models.ChatStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) ChatRoutes(router *mux.Router) {
	router.HandleFunc("/chat", h.HandleGet).Methods("GET")
	router.HandleFunc("/chat/{id}", h.HandleGetByID).Methods("GET")
	router.HandleFunc("/chat", h.HandleCreate).Methods("POST")
	router.HandleFunc("/chat/{id}", h.HandleUpdate).Methods("PUT")
	router.HandleFunc("/chat/{id}", h.HandleDelete).Methods("DELETE")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	chats, err := h.store.GetAllChat()
	if err != nil {
		fmt.Printf("error getting all chats: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to retrieve chats"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, chats)
}

func (h *Handler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil || id <= 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	chat, err := h.store.GetChatByID(id)
	if err != nil || id <= 0 {
		fmt.Printf("error getting chat by id: %v\n", err)
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Chat with ID %d not found", id))
		return
	}

	utils.WriteJSON(w, http.StatusOK, chat)
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var payload models.CreateChatPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		fmt.Printf("error parsing json: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		fmt.Printf("error validating payload: %v\n", err)
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", errors))
		return
	}

	err := h.store.CreateChat(&models.Chat{
		Message: payload.Message,
	})
	if err != nil {
		fmt.Printf("error create chat: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, fmt.Sprintf("Create chat %s successfully", payload.Message))
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var payload models.UpdateChatPayload
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

	chat, err := h.store.GetChatByID(payload.ID)
	if err != nil {
		fmt.Printf("error get chat by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("chat id %d not found"))
		return
	}

	if chat == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Chat with ID %d does not exist", payload.ID))
		return
	}

	if payload.Message == "" {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("New Chat Name Cannot Be Empty!"))
		return
	}

	err = h.store.UpdateChat(&models.Chat{
		ID:      payload.ID,
		Message: payload.Message,
	})
	if err != nil {
		fmt.Printf("error update chat: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Update chat %s successfully", chat.Message))
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid ID parameter"))
		return
	}

	chat, err := h.store.GetChatByID(id)
	if err != nil {
		fmt.Printf("error get chat by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("chat id %d not found"))
		return
	}

	err = h.store.DeleteChat(id)
	if err != nil {
		fmt.Printf("error create user: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Delete chat %s successfully", chat.Message))
}
