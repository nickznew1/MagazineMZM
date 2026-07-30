package service

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/model"
	"github.com/nickznew1/MagazineMZM/backend/internal/domain/usecase"
	"net/http"
)

type ApplicationService struct {
	useCase usecase.ApplicationUseCase
}

func NewApplicationService(c *usecase.ApplicationUseCase) *ApplicationService {
	return &ApplicationService{
		useCase: *c}
}

func (h *ApplicationService) CreateApplication(w http.ResponseWriter, r *http.Request) {
	var input model.Application

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверные данные")
		return
	}
	newApplication, err := h.useCase.CreateApplication(input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"application_id": newApplication,
	})
}

func (h *ApplicationService) GetApplication(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	userIdJWT := r.Context().Value("user_id")
	application, err := h.useCase.GetApplication(idStr, userIdJWT.(string))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	respondWithJSON(w, http.StatusOK, application)
}

func (h *ApplicationService) GetAllApplicationsForUser(w http.ResponseWriter, r *http.Request) {
	userIdJWT := r.Context().Value("user_id")
	applications, err := h.useCase.GetAllApplicationsForUser(userIdJWT.(string))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	respondWithJSON(w, http.StatusOK, applications)
}

func (h *ApplicationService) GetAllApplicationsForAdmin(w http.ResponseWriter, r *http.Request) {
	applications, err := h.useCase.GetAllApplicationsForAdmin()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	respondWithJSON(w, http.StatusOK, applications)
}

func (h *ApplicationService) SetApplicationStatus(w http.ResponseWriter, r *http.Request) {
	var input model.Application

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Неверные данные")
		return
	}
	newApplication, err := h.useCase.SetNewApplicationStatus(input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "oshibka")
		return
	}
	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
		"application_status": newApplication,
	})
}

func (h *ApplicationService) GetApplicationForAdmin(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	itemId, err := h.useCase.GetApplicationForAdmin(idStr)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "id doesnt find")
		return
	}
	RespondWithJSON(w, http.StatusCreated, itemId)
}
