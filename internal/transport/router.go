package transport

import "github.com/gorilla/mux"

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) GetApiRoutes() *mux.Router {
	router := mux.NewRouter()
	// ..
	return router
}
