package handler

import (
	mid "github.com/go-chi/chi/middleware"
	"github.com/gorilla/mux"
	"github.com/namrahov/klawpse/config"
	"github.com/namrahov/klawpse/middleware"
	"github.com/namrahov/klawpse/service"
	"net/http"
)

type bracketsHandler struct {
	BracketsService service.IBracketsService
}

func BracketsHandler(router *mux.Router) *mux.Router {
	router.Use(mid.Recoverer)
	router.Use(middleware.RequestParamsMiddleware)

	h := &bracketsHandler{
		BracketsService: &service.BracketsService{},
	}

	router.HandleFunc(config.RootPath+"/detect", h.detectBracketsType).Methods("POST")

	return router
}

func (h *bracketsHandler) detectBracketsType(w http.ResponseWriter, r *http.Request) {

	filePath := r.URL.Query().Get("filePath")

	err := h.BracketsService.DetectBracketsType(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
