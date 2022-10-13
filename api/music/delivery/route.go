package route

import (
	"test-fullstack/packages/manager"

	musicRoute "test-fullstack/api/music/delivery/route"

	"github.com/gorilla/mux"
)

func NewRoutes(r *mux.Router, mgr manager.Manager) {
	api := r.PathPrefix("/musics").Subrouter()

	musicRoute.NewMusicRoute(mgr, api)
}
