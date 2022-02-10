package handler

import (
	mid "github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	"github.com/namrahov/klawpse/config"
	"github.com/namrahov/klawpse/middleware"
	"github.com/namrahov/klawpse/service"
	"net/http"
)

type palindromeHandler struct {
	PalindromeService service.IBracketsService
}

func PalindromesHandler(router *mux.Router) *mux.Router {
	router.Use(mid.Recoverer)
	router.Use(middleware.RequestParamsMiddleware)

	h := &palindromeHandler{
		PalindromeService: &service.PalindromeService{},
	}

	router.HandleFunc(config.RootPath+"/detect", h.detectPalindromeOfNumber).Methods("POST")

	return router
}

func (h *palindromeHandler) detectPalindromeOfNumber(w http.ResponseWriter, r *http.Request) {

	filePath := r.URL.Query().Get("filePath")

	err := h.PalindromeService.DetectPalindromeOfNumber(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
