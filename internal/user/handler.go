package user

import (
	"net/http"
	"rest_api/internal/apperror"
	"rest_api/internal/handlers"
	"rest_api/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

var _ handlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, apperror.Middleware(h.GetList))
	router.HandlerFunc(http.MethodPost, usersURL, apperror.Middleware(h.CreateUser))
	router.HandlerFunc(http.MethodGet, userURL, apperror.Middleware(h.GetUserByUUID))
	router.HandlerFunc(http.MethodPut, userURL, apperror.Middleware(h.FullyUpdateUser))
	router.HandlerFunc(http.MethodPatch, userURL, apperror.Middleware(h.PartiallyUpdateUser))
	router.HandlerFunc(http.MethodDelete, userURL, apperror.Middleware(h.DeleteUser))

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error{
	w.WriteHeader(200)
	w.Write([]byte("this is list of users"))

	return nil
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) error{
	w.WriteHeader(201)
	w.Write([]byte("create"))

	return nil
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request) error{
	w.WriteHeader(200)
	w.Write([]byte("get user"))

	return nil
}

func (h *handler) FullyUpdateUser(w http.ResponseWriter, r *http.Request) error{
	w.WriteHeader(204)
	w.Write([]byte("fully update user"))

	return nil
}

func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request) error{
	w.WriteHeader(204)
	w.Write([]byte("partially update user"))

	return nil
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) error{
	w.WriteHeader(204)
	w.Write([]byte("delete"))

	return nil
}
