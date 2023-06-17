package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	body, err := ioutil.ReadAll(r.Body)
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

	//// get post parameters
	//err := r.ParseForm()
	//if err != nil {
	//	fmt.Println("ParseForm Error: ", err)
	//}
	//
	//name := r.PostFormValue("name")
	//age := r.PostFormValue("age")
	//fmt.Fprintf(w, "Name: %s\nAge: %s\n", name, age)
	//
	//luteEngine := lute.New() // 默认已经启用 GFM 支持以及中文语境优化
	//html := luteEngine.MarkdownStr("terwer", "**Lute** - A structured markdown engine.")
	//
	//ret := gulu.Ret.NewResult()
	//ret.Data = html
	//
	//// Marshal
	//output, err := sonic.Marshal(&ret)
	//if err != nil {
	//	ret.Msg = err.Error()
	//}
	//fmt.Fprintf(w, string(output))
}
