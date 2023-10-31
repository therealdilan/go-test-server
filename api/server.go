package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Game struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Rating string    `json:"rating"`
}

type Server struct {
	*mux.Router

	gamesList []Game
}

func NewServer() *Server {
	server := &Server{
		Router:    mux.NewRouter(),
		gamesList: []Game{},
	}

	return server
}

func (server *Server) SetupRoutes() {
	server.HandleFunc("/games", server.addNewGame()).Methods("POST")
	server.HandleFunc("/games", server.showGamesList()).Methods("GET")
}

func (server *Server) addNewGame() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Game
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		i.ID = uuid.New()
		server.gamesList = append(server.gamesList, i)

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (server *Server) showGamesList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(server.gamesList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
