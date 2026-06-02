package handlers

import (
	"026_rest_api/internal/usecases"
	"fmt"
	"log/slog"
	"net/http"
)

type Handlers struct {
	useCases *usecases.UseCases
}

func New(useCases *usecases.UseCases) *Handlers {
	return &Handlers{useCases: useCases}
}

func (h Handlers) Listen(port int) error {
	h.registerUserEndpoints()

	slog.Info("listening on", "port", port)

	return http.ListenAndServe(
		fmt.Sprintf(":%v", port),
		nil,
	)
}