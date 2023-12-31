package testgrp

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

// Handlers manages the set of check enpoints.
type Handlers struct {
	Log *zap.SugaredLogger
}

// Test handler is for development
func (h Handlers) Test(w http.ResponseWriter, r *http.Request) {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	json.NewEncoder(w).Encode(status)
	h.Log.Infow("readiness", "statusCode", http.StatusOK, "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)
}
