package utils

import "mime"

func MediaTypeTextHTML() string {
	return mime.FormatMediaType("text/html", map[string]string{"charset": "utf-8"})
}
