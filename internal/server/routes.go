package server

import (
	"net/http"

	"github.com/Evilcmd/Hackup-backend/internal/apis"
	"github.com/Evilcmd/Hackup-backend/internal/middleware"
)

func GetRouter() http.Handler {
	router := http.NewServeMux()

	generalRouter := generalRoutes()
	adminRouter := adminRoutes()

	router.Handle("/", generalRouter)
	router.Handle("/admin/", http.StripPrefix("/admin", middleware.DummyAuth(adminRouter)))

	return middleware.Cors(router)
}

func generalRoutes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", apis.Root)
	router.HandleFunc("GET /health", apis.CheckHealth)
	router.HandleFunc("GET /err", apis.ErrCheck)

	return router
}

func adminRoutes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /", apis.AdminCheck)

	return router
}
