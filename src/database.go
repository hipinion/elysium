package elysium

import (
	"strings"
)

func compileWheres(wheres []string) string {
	whereS := ""
	if len(wheres) > 0 {
		whereS += " WHERE "
		whereS += strings.Join(wheres, " AND ")
	} else {
		return ""
	}
	return whereS
}
