package core

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rakyll/openai-go"
	"github.com/rakyll/openai-go/chat"
)

type ServerCore struct {
}

func ServerCore() *ServerCore {
	return &ServerCore{}
}

// Session is a session created to communicate with OpenAI.
type Session struct {
	// OrganizationID is the ID optionally to be included as
	// a header to requests made from this session.
	// This field must be set before session is used.
	OrganizationID string

	// HTTPClient providing a custom HTTP client.
	// This field must be set before session is used.
	HTTPClient *http.Client

	apiKey string
}

func NewSession(apiKey string) *Session {
	return &Session{
		apiKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (h *ServerCore) SigninQuery(ctx context.Context, query string) (string, error) {
	// return blank if the query is empty
	if query == "" {
		return "", nil
	}

	return "response", nil
}
