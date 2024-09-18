package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandleGetCommits handles parsing input to pass to the GetCommits operation and sends responses back to the client
func (h *APIHandler) HandleGetCommits(w http.ResponseWriter, r *http.Request) {
	var err error
	response, err := h.GetCommits(r.Context())
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetCommits returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetCommits was unable to send it's response, err: %s", err)
	}
}

// HandleHandleWebhook handles parsing input to pass to the HandleWebhook operation and sends responses back to the client
func (h *APIHandler) HandleHandleWebhook(w http.ResponseWriter, r *http.Request) {
	var err error
	reqBody := GitHubWebhook{}
	decoder := json.NewDecoder(r.Body)
	//decoder.DisallowUnknownFields()
	if err := decoder.Decode(&reqBody); err != nil {
		ErrorResponseWithMsg(http.StatusBadRequest, "request body was not able to be parsed successfully 'GitHubWebhook'", w)
		return
	}
	if err := reqBody.Validate(); err != nil {
		errMsg := fmt.Errorf("request body was parsed successfully but failed validation, err: %w", err)
		ErrorResponseWithMsg(http.StatusBadRequest, errMsg.Error(), w)
		return
	}

	response, err := h.HandleWebhook(r.Context(), reqBody)
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("HandleWebhook returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("HandleWebhook was unable to send it's response, err: %s", err)
	}
}
