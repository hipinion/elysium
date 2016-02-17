package elysium

import (
	"log"
	"regexp"
	"strings"
)

func nl2br(t string) string {
	t = strings.Replace(t, "\n", "<br />", -1)
	return t
}
func parseText(text string) string {
	qr := regexp.MustCompile("\\[quote\\=\"(.*?)\"\\:(.*?)\\](.*?)\\[\\/quote\\:.*\\]")
	log.Println(qr)
	text = qr.ReplaceAllString(text, "<div class=\"quote\">$3</div>")
	return text
}
