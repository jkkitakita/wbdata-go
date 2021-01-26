package wbdata

import (
	"net/http"
)

func addFootNoteParams(req *http.Request) {
	params := req.URL.Query()
	params.Set(`footnote`, `y`)
	req.URL.RawQuery = params.Encode()
}
