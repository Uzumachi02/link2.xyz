package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Uzumachi02/link2.xyz/routes"
	"github.com/Uzumachi02/link2.xyz/utils"
)

const PORT = ":8080"

func main() {
	fmt.Printf("Listening Port %s\n", PORT)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
