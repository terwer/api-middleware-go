package handler

import (
	"fmt"
	"github.com/88250/gulu"
	"github.com/88250/lute"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
)

// RequestBody {"md":"# test"}
type RequestBody struct {
	Md string `json:"md"`
}

func handleMarkdown(w http.ResponseWriter, r *http.Request) {
	ret := gulu.Ret.NewResult()

	if r.Method != http.MethodPost {
		ret.Code = http.StatusMethodNotAllowed
		ret.Msg = "Invalid request method, POST only"
		resultBytes, _ := sonic.Marshal(ret)
		fmt.Fprintf(w, "%s", resultBytes)
		return
	}

	// Get post parameters
	maxBodySize := int64(50 * 1024 * 1024)
	body, err := io.ReadAll(io.LimitReader(r.Body, maxBodySize))
	if err != nil {
		ret.Msg = "Error reading request body"
		resultBytes, _ := sonic.Marshal(ret)
		fmt.Fprintf(w, "%s", resultBytes)
		return
	}

	var requestBody RequestBody
	err = sonic.Unmarshal(body, &requestBody)
	if err != nil {
		ret.Msg = "Error parsing JSON request body"
		resultBytes, _ := sonic.Marshal(ret)
		fmt.Fprintf(w, "%s", resultBytes)
		return
	}

	luteEngine := lute.New()
	html := luteEngine.MarkdownStr("terwer", requestBody.Md)

	ret.Data = html

	// Marshal
	output, err := sonic.Marshal(ret)
	if err != nil {
		ret.Msg = err.Error()
		resultBytes, _ := sonic.Marshal(ret)
		fmt.Fprintf(w, "%s", resultBytes)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(ret.Code)
	fmt.Fprintf(w, "%s", output)
}

func handleUnknownType(w http.ResponseWriter, r *http.Request) {
	ret := gulu.Ret.NewResult()
	ret.Code = http.StatusBadRequest
	ret.Msg = "Unknown type parameter"
	resultBytes, _ := sonic.Marshal(ret)
	fmt.Fprintf(w, "%s", resultBytes)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Get type parameter
	t := r.URL.Query().Get("type")
	switch t {
	case "md":
		handleMarkdown(w, r)
	default:
		handleUnknownType(w, r)
	}
}
