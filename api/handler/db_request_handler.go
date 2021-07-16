package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bino1490/case-study/api/response"
	"github.com/bino1490/case-study/pkg/entity"
	"github.com/bino1490/case-study/pkg/logger"
	"github.com/bino1490/case-study/pkg/service"
)

// DBHandler the Handler Layer for future business logic enhancement
// for now just redirecting to Service Layer
func DBHandler(svc service.DBService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Logger.Debug("Entering handler.InMemReqHandler() ...")
		DateMonthFormatConst := "2006-01-02"
		if r.Method != http.MethodPost {
			WriteResponseData(w, "Method Not allowed", http.StatusMethodNotAllowed)
			return
		}
		bodyBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			WriteResponseData(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ct := r.Header.Get("content-type")
		if ct != "application/json" {
			WriteResponseData(w, "Need content-type 'application/json", http.StatusUnsupportedMediaType)
			return
		}

		var memReq entity.DBRequest
		err = json.Unmarshal(bodyBytes, &memReq)
		_, sterr := time.Parse(DateMonthFormatConst, memReq.StartDate)
		_, eterr := time.Parse(DateMonthFormatConst, memReq.EndDate)
		if err != nil || sterr != nil || eterr != nil {
			WriteResponseData(w, "Bad Request or invalid start/end Time", http.StatusBadRequest)
			return
		}

		records, recErr := svc.GetDBRecords(memReq)
		if recErr != nil {
			WriteResponseData(w, "Error Fetching Data", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(BuildSuccessRespBody(records))
	})

}

func WriteResponseData(w http.ResponseWriter, msg string, statusCode int) {
	logger.BootstrapLogger.Error(msg)
	w.WriteHeader(statusCode)
	w.Write(BuildFailureRespBody(msg))
}

func BuildFailureRespBody(msg string) []byte {
	logger.Logger.Debug("Entering handler.BuildFailureRespBody() ...")
	var records []map[string]interface{}
	if msg != "" {
		msg = " - " + msg
	}
	res := response.Response{}
	res.Code = "-1"
	res.Message = "Failure" + msg
	res.Data = records
	resStr, _ := json.Marshal(res)
	return resStr
}

func BuildSuccessRespBody(records []entity.DBRecord) []byte {

	logger.Logger.Debug("Entering handler.buildSuccessRespBody() ...")
	res := response.Response{}
	res.Code = "0"
	res.Message = "Success"
	res.Data = records
	resStr, _ := json.Marshal(res)
	return resStr
}
