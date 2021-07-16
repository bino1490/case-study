package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bino1490/case-study/api/response"
	"github.com/bino1490/case-study/pkg/constant"
	"github.com/bino1490/case-study/pkg/entity"
	"github.com/bino1490/case-study/pkg/logger"
	"github.com/bino1490/case-study/pkg/service"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// getScheduleErrCode contains the list of error codes that can be possibly
// returned by this handler.
var getScheduleErrCode = map[error]string{
	entity.ErrInvalidAccessToken: "4" + constant.ServiceId +
		constant.GetScheduleByChannelId + "01",

	entity.ErrInvalidInputItemId: "4" + constant.ServiceId +
		constant.GetScheduleByChannelId + "02",

	entity.ErrItemNotFound: "4" + constant.ServiceId +
		constant.GetScheduleByChannelId + "03",

	entity.ErrDatabaseFailure: "4" + constant.ServiceId +
		constant.GetScheduleByChannelId + "90",

	entity.ErrDefault: "4" + constant.ServiceId +
		constant.GetScheduleByChannelId + "91",
}

// getScheduleErrDesc contains the list of error descriptions that can be
// possibly returned by this handler.
var getScheduleErrDesc = map[error]string{
	entity.ErrInvalidAccessToken: "Invalid X-Authorization",
	entity.ErrInvalidInputItemId: "Invalid or missing item identifier",
	entity.ErrItemNotFound:       "Item not found",
	entity.ErrDatabaseFailure:    "Subsystem failure",
	entity.ErrDefault:            "Unknown failure",
}

func GetScheduleByChannelID(service service.ScheduleService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		trackingId := r.Context().Value(constant.XTrackingId).(string)
		//trackingId := "123"
		var logFields = logrus.Fields{constant.XTrackingId: trackingId}
		logger.LogDebug("Entering handler.GetBookmark() ...", logFields)

		// #1 - Parse request parameters
		vars := mux.Vars(r)
		channelId := vars["channelId"]
		varsqp := r.URL.Query()
		epochtime := varsqp.Get("t")
		// #2 - Validate input
		errors := validateGetSchedule(channelId, epochtime, logFields)
		if errors != nil {
			w.WriteHeader(http.StatusOK)
			response.Write(w, buildGetScheduleFailureRespBody(
				trackingId, errors, logFields), logFields)
			return
		}

		// #3 - Invoke service to execute business
		item, err := service.GetScheduleByID(channelId, epochtime, logFields)
		if err != nil {
			switch err {
			case entity.ErrInvalidAccessToken:
				w.WriteHeader(http.StatusUnauthorized)
			default:
				w.WriteHeader(http.StatusOK)
			}
			response.Write(w, buildGetScheduleFailureRespBody(
				trackingId, append(errors, err), logFields), logFields)
			return
		}

		// #4 - Build and return success response
		w.WriteHeader(http.StatusOK)
		response.Write(w, buildGetScheduleSuccessRespBody(
			trackingId, item, logFields), logFields)
		return
	})

}

// validateGetSchedule validates the input request payload.
// In this case, the input request payload contains itemId
// Returns error for invalid or missing input parameters
func validateGetSchedule(channelId, epochtime string,
	logFields map[string]interface{}) []error {
	logger.LogDebug("Entering handler.validateGetSchedule() ...", logFields)

	var errors []error

	if channelId == "" || epochtime == "" {
		errors = append(errors, entity.ErrInvalidInputItemId)
		logger.LogDebug(
			"Input request validation error: "+
				getScheduleErrCode[entity.ErrInvalidInputItemId]+
				" - "+
				getScheduleErrDesc[entity.ErrInvalidInputItemId], logFields)
	}
	return errors
}

// buildGetScheduleSuccessRespBody formats the success response payload
// by structuring the input parameters to response structure format
// This builds response header by adding API source, success response code,
// success message and tracking id.
// Returns byte[] of the response payload
func buildGetScheduleSuccessRespBody(trackingId string,
	data interface{}, logFields map[string]interface{}) []byte {
	logger.LogDebug(
		"Entering handler.buildGetScheduleSuccessRespBody() ...", logFields)
	res := response.Response{}
	res.Header = &response.Header{
		Source:     constant.GetScheduleSource,
		Code:       constant.SuccessCode,
		Message:    constant.SuccessMessage,
		SystemTime: (time.Now().UnixNano() / 1e6),
		TrackingId: trackingId,
	}
	res.Data = data
	resStr, _ := json.Marshal(res)
	return resStr
}

// buildGetScheduleFailureRespBody formats the failure response payload
// by structuring the input parameters to response structure format
// This builds response header by adding API source, failure response code,
// failure message, list of errors and tracking id.
// It constructs list of errors by getting error code and error
// descriptions corresponding to lookup bookmark API
// Returns byte[] of the response payload
func buildGetScheduleFailureRespBody(
	trackingId string,
	errlist []error, logFields map[string]interface{}) []byte {
	logger.LogDebug(
		"Entering handler.buildGetScheduleFailureRespBody() ...", logFields)
	var errors []response.Error
	for _, err := range errlist {
		errors = append(errors,
			response.Error{
				Code:        getScheduleErrCode[err],
				Description: getScheduleErrDesc[err],
			})
	}
	res := response.Response{}
	res.Header = &response.Header{
		Source:     constant.GetScheduleSource,
		Code:       constant.FailureCode,
		Message:    constant.FailureMessage,
		SystemTime: (time.Now().UnixNano() / 1e6),
		Errors:     errors,
		TrackingId: trackingId,
	}
	resStr, _ := json.Marshal(res)
	return resStr
}

// logGetBookmarkInputValidationError takes care of logging Get Bookmark
// by identifier request related input validation errors
func logGetBookmarkInputValidationError(err error,
	logFields map[string]interface{}) {
	logger.LogDebug(
		"Input request validation error: "+
			getScheduleErrCode[err]+" - "+getScheduleErrDesc[err],
		logFields)
}
