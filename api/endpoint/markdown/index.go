package markdown

import (
	"fmt"
	"github.com/88250/gulu"
	"github.com/88250/lute"
	"github.com/bytedance/sonic"
	"io"
	"net/http"
)

type ToType int

const (
	HTML ToType = iota
	DOM  ToType = iota
	MD   ToType = iota
	TEXT ToType = iota
)

// RequestBody {"md":"# test", "to": "html"}
type RequestBody struct {
	Md string `json:"md"`
	To ToType `json:"to"`
}

func HandleMarkdown(params *RequestBody) (*gulu.Result, error) {
	ret := gulu.Ret.NewResult()

	// 默认已经启用 GFM 支持以及中文语境优化
	luteEngine := lute.New()

	var html string
	switch params.To {
	case HTML:
		html = luteEngine.MarkdownStr("terwer", params.Md)
	case DOM:
		html = luteEngine.HTML2BlockDOM(params.Md)
	case MD:
		html = luteEngine.HTML2Md(params.Md)
	case TEXT:
		html = luteEngine.HTML2Text(params.Md)
	default:
		html = luteEngine.MarkdownStr("terwer", params.Md)
	}

	ret.Data = html
	return ret, nil
}

func HandleMarkdownEndpoint(w http.ResponseWriter, r *http.Request) {
	ret := gulu.Ret.NewResult()

	if r.Method != http.MethodPost {
		ret.Code = http.StatusMethodNotAllowed
		ret.Msg = "Invalid request method, only POST requests are accepted"
		resultBytes, _ := sonic.Marshal(ret)
		fmt.Fprintf(w, "%s", resultBytes)
		return
	}

	// Get POST parameters
	maxBodySize := int64(50 * 1024 * 1024)
	body, err := io.ReadAll(io.LimitReader(r.Body, maxBodySize))
	if err != nil {
		ret.Msg = "Error occurred while reading request body"
		resultBytes, _ := sonic.Marshal(ret)
		fmt.Fprintf(w, "%s", resultBytes)
		return
	}

	var params RequestBody
	err = sonic.Unmarshal(body, &params)
	if err != nil {
		ret.Msg = "Error occurred while parsing JSON request body"
		resultBytes, _ := sonic.Marshal(ret)
		fmt.Fprintf(w, "%s", resultBytes)
		return
	}

	result, err := HandleMarkdown(&params)
	if err != nil {
		ret.Msg = "Error occurred while processing Markdown"
		resultBytes, _ := sonic.Marshal(ret)
		fmt.Fprintf(w, "%s", resultBytes)
		return
	}

	resultBytes, _ := sonic.Marshal(result)
	fmt.Fprintf(w, "%s", resultBytes)
}
