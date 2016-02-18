package elysium

import (
	"github.com/frustra/bbcode"
	"log"
	_ "regexp"
	"strings"
)

func nl2br(t string) string {
	t = strings.Replace(t, "\n", "<br />", -1)
	return t
}
func parseText(text string) string {
	log.Println("")
	compiler := bbcode.NewCompiler(true, true)
	//text = qr.ReplaceAllString(text, "<div class=\"quote\">$3</div>")
	text = compiler.Compile(text)
	return text
}
