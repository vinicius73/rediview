package server

import (
	"context"
	"github.com/go-chi/httplog"
	"github.com/rs/zerolog"
	"github.com/vinicius73/rediview/pkg/support/httputil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func init() {
	httplog.DefaultOptions.Concise = true
}

func factoryRouter(logger zerolog.Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Use(httplog.RequestLogger(logger))
	r.Use(middleware.Heartbeat("/.heartbeat"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	return r
}

func Start(ctx context.Context, conf Config) error {
	logger := zerolog.Ctx(ctx)
	logger.Info().Msgf("Starting server on %s", conf.Addr)

	r := factoryRouter(*logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		httputil.JSON(w, http.StatusOK, map[string]string{"ok": "ok"})
	})

	return http.ListenAndServe(conf.Addr, r)
}
