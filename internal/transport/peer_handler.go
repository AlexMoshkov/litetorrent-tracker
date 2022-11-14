package transport

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"litetorrent-tracker/internal/dto"
	"litetorrent-tracker/internal/errors"
	"litetorrent-tracker/internal/repositories"
	"litetorrent-tracker/pkg/request_validation"
	"litetorrent-tracker/pkg/responses"
	"net/http"
)

type PeerHandler struct {
	repo *repositories.Repository
}

func NewHandler(repo *repositories.Repository) *PeerHandler {
	return &PeerHandler{repo: repo}
}

func (h *PeerHandler) CreatePeer(w http.ResponseWriter, r *http.Request) {
	data, err := request_validation.DecodeAndValidate[dto.PeerIn](r.Body)
	if err != nil {
		responses.ResponseError(w, 400, err)
		return
	}
	id, err := h.repo.CreatePeer(data.Address.String())
	if err != nil {
		responses.ResponseServerError(w, err)
		return
	}
	responses.ResponseJSON(w, http.StatusOK, dto.PeerOut{Id: *id})
}

func (h *PeerHandler) DeletePeer(w http.ResponseWriter, r *http.Request) {
	peerId, err := uuid.Parse(mux.Vars(r)["peerId"])
	if err != nil {
		responses.ResponseServerError(w, err)
		return
	}
	exists, err := h.repo.IsPeerExist(peerId)
	if err != nil {
		responses.ResponseServerError(w, err)
		return
	}
	if !exists {
		responses.ResponseError(w, 404, errors.NotFoundError)
		return
	}
	if err := h.repo.DeletePeer(peerId); err != nil {
		responses.ResponseServerError(w, err)
		return
	}
	responses.ResponseOK(w)
}

func (h *PeerHandler) UpdateDistributedFiles(w http.ResponseWriter, r *http.Request) {
	peerId, err := uuid.Parse(mux.Vars(r)["peerId"])
	if err != nil {
		responses.ResponseServerError(w, err)
		return
	}
	data, err := request_validation.DecodeAndValidate[dto.FilesIn](r.Body)
	if err != nil {
		responses.ResponseError(w, 400, err)
		return
	}
	exists, err := h.repo.IsPeerExist(peerId)
	if err != nil {
		responses.ResponseServerError(w, err)
		return
	}
	if !exists {
		responses.ResponseError(w, 404, errors.NotFoundError)
		return
	}
	if err := h.repo.UpdateDistributedFiles(peerId, data.DistributionFiles); err != nil {
		responses.ResponseServerError(w, err)
		return
	}
	responses.ResponseOK(w)
}

func (h *PeerHandler) GetPeersAddressesByFile(w http.ResponseWriter, r *http.Request) {
	hash := r.FormValue("field")
	data, err := h.repo.GetPeerAddressesByFile(hash)
	if err != nil {
		responses.ResponseServerError(w, err)
		return
	}
	responses.ResponseJSON(w, 200, data)
}
