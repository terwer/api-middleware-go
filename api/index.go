package handler

import (
	"github.com/terwer/api-middleware-go/api/endpoint/markdown"
	"github.com/terwer/api-middleware-go/api/endpoint/unkonwn"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("type")
	switch t {
	case "md":
		markdown.HandleMarkdownEndpoint(w, r)
	default:
		unkonwn.HandleUnknownType(w, r)
	}
}
