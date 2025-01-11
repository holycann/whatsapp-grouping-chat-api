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
	router.HandleFunc("/user", h.HandleCreate).Methods("POST")
}

func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var payload models.CreateUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid Payload %v", errors))
		return
	}

	_, err := h.store.GetUserByUsername(payload.Username)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("username %s already exists", payload.Username))
		return
	}

	err = h.store.CreateUser(&models.User{
		Username: payload.Username,
		ImageURL: payload.ImageURL,
	})
	if err != nil {
		fmt.Printf("error create user: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var payload models.UpdateUserPayload
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

	if payload.Username == "" && payload.PhoneNumber == "" {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("Username And Phone Number Cannot Be Empty!"))
		return
	}

	err = h.store.UpdateUser(&models.UpdateUserPayload{
		ID:          payload.ID,
		Username:    payload.Username,
		PhoneNumber: payload.PhoneNumber,
		ImageURL:    payload.ImageURL,
	})
	if err != nil {
		fmt.Printf("error update chat: %v\n", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, fmt.Sprintf("Update user %s successfully", u.Username))
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
