package routes

import (
	"net/http"

	"github.com/Uzumachi02/link2.xyz/utils"
)

func homeGetHandler(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "home.html", nil)
}
