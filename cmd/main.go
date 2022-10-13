package main

import (
	"fmt"
	"os"
	musicRoutes "test-fullstack/api/music/delivery"
	"test-fullstack/packages/manager"
	"test-fullstack/packages/server"
	"time"
)

func run() error {
	mgr, err := manager.NewInit()
	if err != nil {
		return err
	}

	tzLocation, err := time.LoadLocation(mgr.GetConfig().AppTz)
	if err != nil {
		return err
	}
	time.Local = tzLocation

	server := server.NewServer(mgr.GetConfig())
	// server.Router.Use(mgr)

	// // routes
	musicRoutes.NewRoutes(server.Router, mgr)
	// reportRoutes.NewRoutes(server.Router, mgr)

	server.RegisterRouter(server.Router)

	return server.ListenAndServe()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}
