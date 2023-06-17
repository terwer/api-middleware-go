package handler

import (
	"encoding/json"
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

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Invalid request method, POST only")
		return
	}

	// get post parameters
	// 50MB 最大大小
	maxBodySize := int64(50 * 1024 * 1024)
	body, err := io.ReadAll(io.LimitReader(r.Body, maxBodySize))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error reading request body")
		return
	}

	var requestBody RequestBody
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing JSON request body")
		return
	}
	fmt.Fprintf(w, "md: %s\n", requestBody.Md)

	luteEngine := lute.New() // 默认已经启用 GFM 支持以及中文语境优化
	html := luteEngine.MarkdownStr("terwer", "**Lute** - A structured markdown engine.")

	ret := gulu.Ret.NewResult()
	ret.Data = html

	// Marshal
	output, err := sonic.Marshal(&ret)
	if err != nil {
		ret.Msg = err.Error()
	}
	fmt.Fprintf(w, string(output))
}
