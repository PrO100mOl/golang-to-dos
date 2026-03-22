package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	stats statsInfc
	hands handsIfc
}

type statsInfc interface {
	getStatsByTasks(http.ResponseWriter, *http.Request)
}

type handsIfc interface {
	newUser(http.ResponseWriter, *http.Request)
	getListUsers(http.ResponseWriter, *http.Request)
	getUserByID(http.ResponseWriter, *http.Request)
	updateUserByID(http.ResponseWriter, *http.Request)
	deleteUserByID(http.ResponseWriter, *http.Request)

	newvTask(http.ResponseWriter, *http.Request)
	getListTask(http.ResponseWriter, *http.Request)
	getTaskByID(http.ResponseWriter, *http.Request)
	updateTaskByID(http.ResponseWriter, *http.Request)
	deleteTaskByID(http.ResponseWriter, *http.Request)
}

func NewServer(stats statsInfc,
	hands handsIfc,
) server {
	return server{
		stats: stats,
		hands: hands,
	}
}

func (s server) Start() error {
	router := mux.NewRouter()

	router.Path("/users").Methods("POST").HandlerFunc(s.hands.newUser)
	router.Path("/users").Methods("GET").Queries("limit", "{limit}", "offset", "{offset}").HandlerFunc(s.hands.getListUsers)
	router.Path("/users/{id}").Methods("GET").HandlerFunc(s.hands.getUserByID)
	router.Path("/users/{id}").Methods("PATCH").HandlerFunc(s.hands.updateUserByID)
	router.Path("/users/{id}").Methods("DELETE").HandlerFunc(s.hands.deleteUserByID)

	router.Path("/tasks").Methods("POST").HandlerFunc(s.hands.newvTask)
	router.Path("/tasks").Methods("GET").Queries("user_id", "{user_id}", "limit", "{limit}", "offset", "{offset}").HandlerFunc(s.hands.getListTask)
	router.Path("/tasks/{id}").Methods("GET").HandlerFunc(s.hands.getTaskByID)
	router.Path("/tasks/{id}").Methods("PATCH").HandlerFunc(s.hands.updateTaskByID)
	router.Path("/tasks/{id}").Methods("DELETE").HandlerFunc(s.hands.deleteTaskByID)

	router.Path("/stats").Methods("GET").Queries("user_id", "{user_id}")

	err := http.ListenAndServe(":8080", router)
	return err
}
