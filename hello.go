package main

import (
	"fmt"

	"github.com/88250/lute"
)

func main() {
	luteEngine := lute.New() // 默认已经启用 GFM 支持以及中文语境优化
	html := luteEngine.MarkdownStr("demo", "**Lute** - A structured markdown engine.")
	fmt.Println(html)
	// <p><strong>Lute</strong> - A structured Markdown engine.</p>
}
