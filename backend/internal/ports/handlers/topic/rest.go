package topic

import (
	"context"
	"encoding/json"
	"github.com/aghex70/daps/internal/core/usecases/topic"
	"github.com/aghex70/daps/internal/ports/handlers"
	"github.com/aghex70/daps/internal/ports/requests/topic"
	"github.com/aghex70/daps/internal/ports/responses"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	CreateTopicUseCase topic.CreateTopicUseCase
	DeleteTopicUseCase topic.DeleteTopicUseCase
	GetTopicUseCase    topic.GetTopicUseCase
	ListTopicsUseCase  topic.ListTopicsUseCase
	UpdateTopicUseCase topic.UpdateTopicUseCase
	logger             *log.Logger
}

func (h Handler) HandleTopics(w http.ResponseWriter, r *http.Request) {
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
	payload := requests.CreateTopicRequest{}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	t, err := h.CreateTopicUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(responses.CreateEntityResponse{ID: t.ID})
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

	filters := make(map[string]interface{})
	for k, v := range r.URL.Query() {
		if len(v) == 1 {
			// If there is only one value, use it directly
			filters[k] = v[0]
		} else if len(v) > 1 {
			// If there are multiple values, use a slice
			filters[k] = v
		}
	}
	topics, err := h.ListTopicsUseCase.Execute(context.TODO(), &filters, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(topics)
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
	return
}

func (h Handler) HandleTopic(w http.ResponseWriter, r *http.Request) {
	// Get topic id & action (if present) from request URI
	path := strings.Split(r.RequestURI, handlers.TOPIC_STRING)[1]
	t := strings.Split(path, "/")[0]
	topicID, err := strconv.Atoi(t)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.Get(w, r, uint(topicID))
	case http.MethodDelete:
		h.Delete(w, r, uint(topicID))
	case http.MethodPut:
		h.Update(w, r, uint(topicID))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.GetTopicRequest{TopicID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	t, err := h.GetTopicUseCase.Execute(context.TODO(), payload, userID)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(t)
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
	payload := requests.DeleteTopicRequest{TopicID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.DeleteTopicUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) Update(w http.ResponseWriter, r *http.Request, id uint) {
	payload := requests.UpdateTopicRequest{TopicID: id}
	if err := handlers.ValidateRequest(r, &payload); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	userID, err := handlers.RetrieveJWTClaims(r, nil)
	if err != nil {
		handlers.ThrowError(err, http.StatusUnauthorized, w)
		return
	}

	if err = h.UpdateTopicUseCase.Execute(context.TODO(), payload, userID); err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func NewTopicHandler(
	createTopicUseCase *topic.CreateTopicUseCase,
	deleteTopicUseCase *topic.DeleteTopicUseCase,
	getTopicUseCase *topic.GetTopicUseCase,
	listTopicsUseCase *topic.ListTopicsUseCase,
	updateTopicUseCase *topic.UpdateTopicUseCase,
	logger *log.Logger,
) *Handler {
	return &Handler{
		CreateTopicUseCase: *createTopicUseCase,
		DeleteTopicUseCase: *deleteTopicUseCase,
		GetTopicUseCase:    *getTopicUseCase,
		ListTopicsUseCase:  *listTopicsUseCase,
		UpdateTopicUseCase: *updateTopicUseCase,
		logger:             logger,
	}
}
