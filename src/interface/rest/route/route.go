package route

import (
	layananHandler "auth-skm/src/interface/rest/handlers/layanan"
	userHandler "auth-skm/src/interface/rest/handlers/user"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserAppRouter(uh userHandler.UserHandlerInterface) http.Handler {
	r := chi.NewRouter()

	// not using middleware yet
	// r.Use(middleware.CheckAPWebHeader)

	r.Mount("/", UserRouter(uh))

	return r
}

func LayananAppRouter(lh layananHandler.LayananHandlerInterface) http.Handler {
	r := chi.NewRouter()

	// not using middleware yet
	// r.Use(middleware.CheckAPWebHeader)

	r.Mount("/", LayananRouter(lh))

	return r
}
