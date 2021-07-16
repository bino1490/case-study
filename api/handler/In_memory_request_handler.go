package handler

import (
	"net/http"

	"github.com/bino1490/case-study/pkg/logger"
	"github.com/bino1490/case-study/pkg/service"
)

func InMemReqHandler(inMemSvc *service.MemHandlers) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		logger.Logger.Debug("Entering handler.InMemReqHandler() ...")

		inMemSvc.InMemGetPOST(w, r)
		return
	})

}
