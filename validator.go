package next

import (
	"net/http"

	"next/internal/binding"
)

// ShouldBindJSON .
func ShouldBindJSON(req *http.Request, obj interface{}) error {
	return binding.JSON.Bind(req, obj)
}
