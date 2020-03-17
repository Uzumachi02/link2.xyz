package main

import (
	"fmt"

	"github.com/Uzumachi02/link2.xyz/routes"
)

const PORT = ":8080"

func main() {
	r := routes.New()

	fmt.Printf("Listening Port %s\n", PORT)
	r.Logger.Fatal(r.Start(PORT))
}
