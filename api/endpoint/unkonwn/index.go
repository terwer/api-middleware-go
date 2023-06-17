package unkonwn

import (
	"fmt"
	"github.com/88250/gulu"
	"github.com/bytedance/sonic"
	"net/http"
)

func HandleUnknownType(w http.ResponseWriter, r *http.Request) {
	ret := gulu.Ret.NewResult()
	ret.Code = http.StatusBadRequest
	ret.Msg = fmt.Sprintf("Unknown type parameter: %s", r.URL.Query().Get("type"))
	resultBytes, _ := sonic.Marshal(ret)
	fmt.Fprintf(w, "%s", resultBytes)
}
