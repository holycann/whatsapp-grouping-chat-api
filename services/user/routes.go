package user

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
	store models.UserStore
}

func NewHandler(store models.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) UserRoutes(router *mux.Router) {
	router.HandleFunc("/user", h.HandleGet).Methods("GET")
	router.HandleFunc("/user/{id}", h.HandleGetByID).Methods("GET")
	router.HandleFunc("/user", h.HandleCreate).Methods("POST")
	router.HandleFunc("/user/{id}", h.HandleUpdate).Methods("PUT")
	router.HandleFunc("/user/{id}", h.HandleDelete).Methods("DELETE")
}

func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	chats, err := h.store.GetAllUser()
	if err != nil {
		fmt.Printf("error getting all user: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to retrieve users"))
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

	chat, err := h.store.GetUserByID(id)
	if err != nil || id <= 0 {
		fmt.Printf("error getting user by id: %v\n", err)
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("User with ID %d not found", id))
		return
	}

	utils.WriteJSON(w, http.StatusOK, chat)
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var payload models.CreateUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v\n", err))
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	_, err := h.store.GetUserByUsername(payload.Name)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Name %s already exists", payload.Name))
		return
	}

	err = h.store.CreateUser(&models.User{
		Name:        payload.Name,
		PhoneNumber: payload.PhoneNumber,
		ImageURL:    payload.ImageURL,
	})
	if err != nil {
		fmt.Printf("error create user: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, fmt.Sprintf("Create user %s successfully", payload.Name))
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var payload models.UpdateUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error parsing json: %v", err))
		return
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
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", err.(validator.ValidationErrors)))
		return
	}

	u, err := h.store.GetUserByID(payload.ID)
	if err != nil {
		fmt.Printf("error get user by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user id %d not found"))
		return
	}

	if u == nil {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("User with ID %d does not exist", payload.ID))
		return
	}

	if payload.Name == "" && payload.PhoneNumber == "" {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Username And Phone Number Cannot Be Empty!"))
		return
	}

	err = h.store.UpdateUser(&models.UpdateUserPayload{
		ID:          payload.ID,
		Name:        payload.Name,
		PhoneNumber: payload.PhoneNumber,
		ImageURL:    payload.ImageURL,
	})
	if err != nil {
		fmt.Printf("error update chat: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Update user %s successfully", u.Name))
}

func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Printf("error get user by id: %v\n", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user id %d not found"))
		return
	}

	err = h.store.DeleteUser(id)
	if err != nil {
		fmt.Printf("error delete user: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Delete user successfully"))
}
