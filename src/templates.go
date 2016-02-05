package elysium

import (
	"html/template"
)

var Templates = template.Must(template.ParseGlob("templates/*.html"))
