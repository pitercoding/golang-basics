package main

import "fmt"

func main() {
	mapa := map[string]int{"Quase Nada": 25, "Racha Cuca": 30}
	for key, value := range mapa {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}