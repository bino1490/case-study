package response

import (
	"net/http"

	"github.com/bino1490/case-study/pkg/logger"
)

// Write writes the HTTP response payload.
func Write(w http.ResponseWriter, data []byte,
	logFields map[string]interface{}) {
	if data != nil {
		_, err := w.Write(data)
		if err != nil {
			logger.LogError(
				"Error in writing to Response writer", logFields)
			logger.LogError(err, logFields)
		}
	}
}
