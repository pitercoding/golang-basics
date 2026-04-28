package main

import (
	"026_rest_api/internal/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	req := models.CreateUserRequest {
		Name: "Racha Cuca",
		Email: "rc@test.com",
	}

	b, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:8080/users", "application/json", bytes.NewReader(b))
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusCreated {
		var responseApi models.ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&responseApi); err != nil {
		panic(err)
	}
		panic(responseApi.Reason)
	}

	var responseApi models.CreateUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseApi); err != nil {
		panic(err)
	}

	fmt.Println("new user created", responseApi)
}
