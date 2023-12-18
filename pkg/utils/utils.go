package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, data interface{}) error {
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, data); err != nil {
			return err
		}
	}
	return nil
}
