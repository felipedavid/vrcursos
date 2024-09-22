package helper

import "net/http"

func MessageResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	res := map[string]any{"status": status, "message": message}
	WriteJSON(w, status, res, nil)
}
