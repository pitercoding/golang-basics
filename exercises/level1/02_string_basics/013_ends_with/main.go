package main

import (
	"fmt"
	"strings"
)

func main() {
	
	result := strings.HasSuffix(
		strings.ToLower("photo.PNG"), 
		strings.ToLower(".png"))

	fmt.Println("Ends with suffix:", result)
}