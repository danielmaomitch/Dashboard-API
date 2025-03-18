package api

import (
	"encoding/json"
	"go-http/server/api/TaskList"
	"net/http"

	"github.com/gorilla/mux"
)

// type Item struct {
// 	ID   uuid.UUID `json:"id"`
// 	Name string    `json:"name"`
// }

func (s *Server) taskRoutes() {
	s.HandleFunc("/task-mnger", s.listTasks()).Methods("GET")
	s.HandleFunc("/task-mnger", s.createTask()).Methods("POST")
	s.HandleFunc("/task-mnger/{user}/{id}", s.removeTask()).Methods("DELETE")
}

func (s *Server) createTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var t TaskList.Task
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id := TaskList.AddRecord(t)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// result := getRecords("twiggs", "Task")
		// taskList := unmarshalTasks(result)

		taskList := TaskList.GetRecord()
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(taskList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id := mux.Vars(r)["id"]
		userID := mux.Vars(r)["user"]
		SK, err := TaskList.DelRecord(userID, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(SK); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}
