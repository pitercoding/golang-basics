package main

import (
	"026_rest_api/internal/handlers"
	"026_rest_api/internal/repositories"
	"026_rest_api/internal/usecases"
)

// REST API: cadastrar e registrar usuários
// handler/controller (recebe requisições HTTP, processa e retorna resposta ) <- usecases/services (lógica) <- repositories (chamadas para BD)
func main() {
	repos := repositories.New()
	useCases := usecases.New(repos)
	h := handlers.New(useCases)

	h.Listen(8080)
}