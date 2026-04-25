package main

import "fmt"

func main() {
	mapa := map[string]int{"Racha Cuca": 25, "Quase Nada": 30}
	for key, value := range mapa {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}