package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// APIHandler is a type to give the api functions below access to a common logger
// any any other shared objects
type APIHandler struct {
	// Zerolog was chosen as the default logger, but you can replace it with any logger of your choice
	logger zerolog.Logger

	// Note: if you need to pass in a client for your database, this would be a good place to include it
}

func NewAPIHandler() *APIHandler {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &APIHandler{logger: logger}
}

func (h *APIHandler) WithLogger(logger zerolog.Logger) *APIHandler {
	h.logger = logger
	return h
}

// This endpoint retrieves a list of commits made to the repository.
// Get a list of commits
func (h *APIHandler) GetCommits(ctx context.Context) (Response, error) {
	// TODO: implement the GetCommits function to return the following responses

	// return NewResponse(200, []Commit, "application/json", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"getCommits operation has not been implemented yet"}, "application/json", nil), nil
}

// This endpoint receives webhook events from GitHub when a commit is made.
// Handle GitHub webhook for commits
func (h *APIHandler) HandleWebhook(ctx context.Context, reqBody GitHubWebhook) (Response, error) {
	fmt.Println(reqBody.Commits)

	return NewResponse(200, WebhookResponse{Message: "Successfully handled webhook!"}, "application/json", nil), nil
}
