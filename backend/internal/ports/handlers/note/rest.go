package note

import (
	"context"
	"encoding/json"
	"github.com/aghex70/daps/internal/core/usecases/note"
	"github.com/aghex70/daps/internal/ports/handlers"
	"github.com/aghex70/daps/internal/ports/requests/note"
	"github.com/aghex70/daps/internal/ports/responses"
	noteResponses "github.com/aghex70/daps/internal/ports/responses/note"

	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	CreateNoteUseCase  note.CreateNoteUseCase
	DeleteNoteUseCase  note.DeleteNoteUseCase
	GetNoteUseCase     note.GetNoteUseCase
	ListNotesUseCase   note.ListNotesUseCase
	ShareNoteUseCase   note.ShareNoteUseCase
	UnshareNoteUseCase note.UnshareNoteUseCase
	UpdateNoteUseCase  note.UpdateNoteUseCase
	logger             *log.Logger
}

func (h Handler) HandleNotes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.List(w, r)
	case http.MethodPost:
		h.Create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	payload := requests.CreateNoteRequest{}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	n, err := h.CreateNoteUseCase.Execute(context.TODO(), userID, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(responses.CreateEntityResponse{ID: n.ID})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
	return
}

func (h Handler) List(w http.ResponseWriter, r *http.Request) {
	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	notes, err := h.ListNotesUseCase.Execute(context.TODO(), nil, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(noteResponses.ListNotesResponse{Notes: notes})
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
	return
}

func (h Handler) HandleNote(w http.ResponseWriter, r *http.Request) {
	// Get note id & action (if present) from request URI
	path := strings.Split(r.RequestURI, handlers.NOTE_STRING)[1]
	n := strings.Split(path, "/")[0]
	noteID, err := strconv.Atoi(n)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	if strings.Contains(r.RequestURI, handlers.SHARE_STRING) {
		h.Share(w, r, uint(noteID))
	} else if strings.Contains(r.RequestURI, handlers.UNSHARE_STRING) {
		h.Unshare(w, r, uint(noteID))
	} else {
		switch r.Method {
		case http.MethodGet:
			h.Get(w, r, uint(noteID))
		case http.MethodDelete:
			h.Delete(w, r, uint(noteID))
		case http.MethodPut:
			h.Update(w, r, uint(noteID))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.GetNoteRequest{NoteID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	n, err := h.GetNoteUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(n)
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
	return
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.DeleteNoteRequest{NoteID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.DeleteNoteUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.UpdateNoteRequest{NoteID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.UpdateNoteUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Share(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.ShareNoteRequest{NoteID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.ShareNoteUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h Handler) Unshare(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.UnshareNoteRequest{NoteID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.UnshareNoteUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NewNoteHandler(
	createNoteUseCase *note.CreateNoteUseCase,
	deleteNoteUseCase *note.DeleteNoteUseCase,
	getNoteUseCase *note.GetNoteUseCase,
	listNotesUseCase *note.ListNotesUseCase,
	shareNoteUseCase *note.ShareNoteUseCase,
	unshareNoteUseCase *note.UnshareNoteUseCase,
	updateNoteUseCase *note.UpdateNoteUseCase,
	logger *log.Logger,
) *Handler {
	return &Handler{
		CreateNoteUseCase:  *createNoteUseCase,
		DeleteNoteUseCase:  *deleteNoteUseCase,
		GetNoteUseCase:     *getNoteUseCase,
		ListNotesUseCase:   *listNotesUseCase,
		ShareNoteUseCase:   *shareNoteUseCase,
		UnshareNoteUseCase: *unshareNoteUseCase,
		UpdateNoteUseCase:  *updateNoteUseCase,
		logger:             logger,
	}
}
