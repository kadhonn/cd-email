package server

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
)

func renderSite(w io.Writer) {
	_, err := w.Write([]byte("<html><head><title>Hey</title></head><body><form accept-charset=\"UTF-8\" method=\"POST\" action=\"/submit/\" enctype=\"multipart/form-data\"><input type=\"text\" name=\"email\" placeholder=\"email\" /><input type=\"password\" name=\"password\" placeholder=\"password\" /><input type=\"file\" name=\"file\" /><input type=\"submit\" value=\"Send\" /></form></body></html>"))
	if err != nil {
		log.Print(err)
	}
}

func renderFavico(w io.Writer) {
	content, _ := base64.StdEncoding.DecodeString("AAABAAEAEBAAAAEAIABoBAAAFgAAACgAAAAQAAAAIAAAAAEAIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOTk5JHl5eaZbW1vvNDQ0/AAAAPQAAADaAAAAewAAABsAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAACdXV1pdzc3P4+Pj7/AAAA/wAAAP8AAAD/AAAA/wAAAP8AAAD1AAAAewAAAAIAAAAAAAAAAAAAAAAAAAACioqK0f7+/v+Hh4f/AAAA/xEREf+Ghob/m5ub/yUlJf8AAAD/AAAA/wAAAP8AAACrAAAAAgAAAAAAAAAAb29vpf39/f//////MTEx/wAAAP8AAAD/Ozs7/35+fv8vLy//AAAA/wAAAP8AAAD/AAAA/wAAAH0AAAAAMDAwJd7e3v7//////////ygoKP8AAAD/AAAA/zs7O/9AQED/AAAA/wAAAP8AAAD/AAAA/wAAAP8AAAD2AAAAHXp6eqX///////////////9nZ2f/AAAA/xAQEP+JiYn/QEBA/wAAAP8AAAD/AAAA/wAAAP8AAAD/AAAA/wAAAHmoqKjn////////////////6enp/xUVFf8AAAD/EBAQ/wQEBP8AAAD/AAAA/wAAAP8AAAD/AAAA/wAAAP8AAADAvb299v/////////////////////i4uL/VlZW/xUVFf8AAAD/AAAA/wAAAP8AAAD/AAAA/wAAAP8AAAD/AAAA4Ly8vPb/////////////////////////////////////8fHx/6ampv8XFxf/AAAA/wAAAP8AAAD/AAAA/wAAAOCnp6fn////////////////////////////////////////////////4+Pj/xISEv8AAAD/AAAA/wAAAP8AAADDenp6qv//////////////////////////3d3d/2dnZ/9oaGj/3d3d//////+VlZX/AAAA/wAAAP8AAAD/AAAAgDAwMCXh4eH+/////////////////////9HR0f+1tbX/t7e3/9HR0f//////1dXV/wAAAP8AAAD/AAAA9gAAAB0AAAAAeXl5qf7+/v/////////////////R0dH/tbW1/7e3t//R0dH//////83Nzf8AAAD/AAAA/wAAAH0AAAAAAAAAAAAAAAKKiorS/v7+////////////5+fn/3Z2dv92dnb/5+fn//////91dXX/AAAA/wAAAKsAAAACAAAAAAAAAAAAAAAAAAAAAnV1daXg4OD+//////////////////////////+/v7//BQUF9QAAAHsAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAOTk5JIODg6ivr6/ovr6+9b6+vvR3d3faBAQEewAAABsAAAAAAAAAAAAAAAAAAAAA+D8AAOAPAADAAwAAgAMAAIABAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAAAAACAAQAAgAMAAMADAADgDwAA+D8AAA==")
	_, err := w.Write(content)
	if err != nil {
		log.Print(err)
	}
}

func renderOk(w io.Writer) {
	_, err := w.Write([]byte("<html><head><title>Ok</title></head><body>Sent!</body></html>"))
	if err != nil {
		log.Print(err)
	}
}

func renderError(w io.Writer, e error) {
	_, err := fmt.Fprintf(w, "<html><head><title>Error</title></head><body>Uh oh! Error: %s </body></html>", e.Error())
	if err != nil {
		log.Print(err)
	}
}
