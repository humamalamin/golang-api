package handler

import (
	"encoding/json"
	"net/http"
	musicDomainEntity "test-fullstack/api/music/domain/entity"
	musicDomainInterface "test-fullstack/api/music/domain/interface"
	musicUsecase "test-fullstack/api/music/usecase"
	jsonHelpers "test-fullstack/helpers/json"
	"test-fullstack/packages/manager"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type Music struct {
	Usecase musicDomainInterface.MusicUsecase
}

func NewMusicHandler(mgr manager.Manager) musicDomainInterface.MusicHandler {
	handler := new(Music)
	handler.Usecase = musicUsecase.NewMusicUsecase(mgr)

	return handler
}

func (h *Music) Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		results, err := h.Usecase.Get(r.Context())
		if err != nil {
			code := "[Handler] Get-1"
			log.Error().Err(err).Msg(code)
			jsonHelpers.ErrorResponse(w, r, false, http.StatusBadRequest, err.Error(), nil)
			return
		}

		jsonHelpers.SuccessResponse(w, r, true, http.StatusOK, "1000", results, "Music List", nil)
	})
}

func (h *Music) GetByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := mux.Vars(r)
		results, err := h.Usecase.GetByID(r.Context(), param["id"])
		if err != nil {
			code := "[Handler] GetByID-1"
			log.Error().Err(err).Msg(code)
			jsonHelpers.ErrorResponse(w, r, false, http.StatusBadRequest, err.Error(), nil)
			return
		}

		jsonHelpers.SuccessResponse(w, r, true, http.StatusOK, "1000", results, "Detail Music", nil)
	})
}

func (h *Music) Store() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request := musicDomainEntity.RequestMusic{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			code := "[Handler] Store-1"
			log.Error().Err(err).Msg(code)
			jsonHelpers.ErrorResponse(w, r, false, http.StatusUnprocessableEntity, err.Error(), nil)
			return
		}

		err := h.Usecase.Store(r.Context(), request)
		if err != nil {
			code := "[Handler] Store-2"
			log.Error().Err(err).Msg(code)
			jsonHelpers.ErrorResponse(w, r, false, http.StatusBadRequest, err.Error(), nil)
			return
		}

		jsonHelpers.SuccessResponse(w, r, true, http.StatusCreated, "1001", nil, "Insert Music", nil)
	})
}

func (h *Music) Update() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := mux.Vars(r)
		request := musicDomainEntity.RequestMusic{}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			code := "[Handler] Store-1"
			log.Error().Err(err).Msg(code)
			jsonHelpers.ErrorResponse(w, r, false, http.StatusUnprocessableEntity, err.Error(), nil)
			return
		}

		err := h.Usecase.Update(r.Context(), request, param["id"])
		if err != nil {
			code := "[Handler] Update-2"
			log.Error().Err(err).Msg(code)
			jsonHelpers.ErrorResponse(w, r, false, http.StatusBadRequest, err.Error(), nil)
			return
		}

		jsonHelpers.SuccessResponse(w, r, true, http.StatusOK, "1000", nil, "Update Music", nil)
	})
}

func (h *Music) Delete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := mux.Vars(r)
		err := h.Usecase.Delete(r.Context(), param["id"])
		if err != nil {
			code := "[Handler] Delete-1"
			log.Error().Err(err).Msg(code)
			jsonHelpers.ErrorResponse(w, r, false, http.StatusBadRequest, err.Error(), nil)
			return
		}

		jsonHelpers.SuccessResponse(w, r, true, http.StatusOK, "1000", nil, "Delete Music", nil)
	})
}
