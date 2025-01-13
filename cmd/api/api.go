package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/holycann/whatsapp-grouping-chat-api/services/chat"
	"github.com/holycann/whatsapp-grouping-chat-api/services/folder"
	"github.com/holycann/whatsapp-grouping-chat-api/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.UserRoutes(subrouter)

	chatStore := chat.NewStore(s.db)
	chatHandler := chat.NewHandler(chatStore)
	chatHandler.ChatRoutes(subrouter)

	folderStore := folder.NewStore(s.db)
	folderHandler := folder.NewHandler(folderStore)
	folderHandler.FolderRoutes(subrouter)

	log.Print("Listening On Port ", s.addr)

	return http.ListenAndServe(s.addr, router)
}
