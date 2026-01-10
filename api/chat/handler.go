package chat

import (
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"super-llm/domain/committee"

	"github.com/cv70/pkgo/llm"
)

// Handler holds dependencies for the chat handler
type Handler struct {
	committee *committee.CommitteeDomain
}

// NewHandler creates a new chat handler
func NewHandler(committee *committee.CommitteeDomain) *Handler {
	return &Handler{
		committee: committee,
	}
}

// ChatCompletions handles the /chat/completions endpoint
func (h *Handler) ChatCompletions(c *gin.Context) {
	// Parse request body
	var req llm.ChatCompletionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Parse headers
	membersHeader := c.GetHeader("X-Members")
	viewsHeader := c.GetHeader("X-Views")

	// Process models from header
	members := parseModels(membersHeader)

	// Process stage from header
	// opinion, review
	var opinion, review bool
	if viewsHeader != "" {
		opinion, review = parseViews(viewsHeader)
	}

	// Log the request
	slog.Info(
		"Received chat completions request",
		slog.Any("members", membersHeader),
		slog.Any("views", viewsHeader),
		slog.Any("messages_count", len(req.Messages)),
	)

	// Process the request using committee and LLM service
	response, err := h.processRequest(c, &req, members, opinion, review)
	if err != nil {
		slog.Error("Failed to process chat completions", slog.Any("err", err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer response.Body.Close()

	// Return response
	if req.Stream {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
	} else {
		c.Writer.Header().Set("Content-Type", "application/json")
	}
	_, err = io.Copy(c.Writer, response.Body)
	if err != nil {
		slog.Error("copy response", slog.Any("err", err))
	}
}

// processRequest processes the chat completion request
func (h *Handler) processRequest(c *gin.Context, req *llm.ChatCompletionRequest, members []string, opinion, review bool) (*http.Response, error) {
	// For simplicity, we'll use the RunCommitteeProcess method directly
	// Since our interface requires a single question, we'll just use the first user message
	result, err := h.committee.RunCommitteeProcess(c, req, members, opinion, review)
	return result, err
}

// parseModels parses comma-separated model names from header
func parseModels(header string) []string {
	// Simple split by comma, trim spaces
	models := []string{}
	for _, model := range strings.Split(header, ",") {
		model = strings.TrimSpace(model)
		if model != "" {
			models = append(models, model)
		}
	}
	return models
}

// parseViews parses view from header
func parseViews(header string) (bool, bool) {
	parts := strings.Split(header, ",")
	opinion := false
	review := false
	for _, part := range parts {
		switch strings.ToLower(strings.TrimSpace(part)) {
		case "opinion":
			opinion = true
		case "review":
			review = true
		}
	}
	return opinion, review
}
