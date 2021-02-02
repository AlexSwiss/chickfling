package controllers

import (
	"net/http"

	"github.com/AlexSwiss/chickfling/api/responses"
)

// Home returns welcome message
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To The Chickfling API")

}
