package musicRoute

import (
	"test-fullstack/packages/manager"

	musicHandler "test-fullstack/api/music/delivery/handler"

	"github.com/gorilla/mux"
)

func NewMusicRoute(mgr manager.Manager, route *mux.Router) {
	MusicHandler := musicHandler.NewMusicHandler(mgr)

	route.Handle("/", MusicHandler.Get()).Methods("GET")
	route.Handle("/", MusicHandler.Store()).Methods("POST")
	route.Handle("/{id}", MusicHandler.GetByID()).Methods("GET")
	route.Handle("/{id}", MusicHandler.Update()).Methods("PUT")
	route.Handle("/{id}", MusicHandler.Delete()).Methods("Delete")
}
